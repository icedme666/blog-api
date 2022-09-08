package article_service

import (
	"github.com/golang/freetype"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"
	"xiamei.guo/blog-api/pkg/file"
	"xiamei.guo/blog-api/pkg/qrcode"
	"xiamei.guo/blog-api/pkg/setting"
)

type ArticlePoster struct {
	PosterName string
	*Article
	Qr *qrcode.QrCode
}

func NewArticlePoster(posterName string, article *Article, qr *qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{
		PosterName: posterName,
		Article:    article,
		Qr:         qr,
	}
}

func GetPosterFlag() string { return "poster" }

//检查合并后图像（指的是存放合并后的海报）是否存在

func (a *ArticlePoster) CheckMergedImage(path string) bool {
	if file.CheckNotExist(path+a.PosterName) == true {
		return false
	}
	return true
}

func (a *ArticlePoster) OpenMergedImage(path string) (*os.File, error) {
	f, err := file.MustOpen(a.PosterName, path)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type ArticlePosterBg struct {
	Name string
	*ArticlePoster
	*Rect
	*Pt
}

type Rect struct {
	Name string
	X0   int
	Y0   int
	X1   int
	Y1   int
}

type Pt struct {
	X int
	Y int
}

func NewArticlePosterBg(name string, ap *ArticlePoster, rect *Rect, pt *Pt) *ArticlePosterBg {
	return &ArticlePosterBg{
		Name:          name,
		ArticlePoster: ap,
		Rect:          rect,
		Pt:            pt,
	}
}

func (a *ArticlePosterBg) Generate() (string, string, error) {
	fullPath := qrcode.GetQrCodeFullPath()
	fileName, path, err := a.Qr.Encode(fullPath)
	if err != nil {
		return "", "", err
	}

	//若不存在，则生成待合并的图像 mergedF
	if !a.CheckMergedImage(path) {
		mergedF, err := a.OpenMergedImage(path)
		if err != nil {
			return "", "", err
		}
		defer mergedF.Close()

		bgF, err := file.MustOpen(a.Name, path) //打开事先存放的背景图 bgF
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()

		qrF, err := file.MustOpen(fileName, path) //打开生成的二维码图像 qrF
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()

		//解码 bgF 和 qrF 返回 image.Image
		bgImage, err := jpeg.Decode(bgF)
		if err != nil {
			return "", "", err
		}
		qrImage, err := jpeg.Decode(qrF)
		if err != nil {
			return "", "", err
		}

		//创建一个新的 RGBA 图像
		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))

		//在 RGBA 图像上绘制 背景图（bgF）
		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		//在已绘制背景图的 RGBA 图像上，在指定 Point 上绘制二维码图像（qrF）
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)
		//将绘制好的 RGBA 图像以 JPEG 4：2：0 基线格式写入合并后的图像文件（mergedF）
		//jpeg.Encode(mergedF, jpg, nil)

		//绘制文字
		err = a.DrawPoster(&DrawText{
			JPG:    jpg,
			Merged: mergedF,

			Title: "First Blog By Go",
			X0:    80,
			Y0:    160,
			Size0: 42,

			SubTitle: "--gxm",
			X1:       320,
			Y1:       220,
			Size1:    36,
		}, "msyhbd.ttc")

		if err != nil {
			return "", "", err
		}
	}
	return fileName, path, nil
}

type DrawText struct {
	JPG    draw.Image
	Merged *os.File

	Title string
	X0    int
	Y0    int
	Size0 float64

	SubTitle string
	X1       int
	Y1       int
	Size1    float64
}

func (a *ArticlePosterBg) DrawPoster(d *DrawText, fontName string) error {
	fontSource := setting.AppSetting.RuntimeRootPath + setting.AppSetting.FontSavePath + fontName
	fontSourceBytes, err := ioutil.ReadFile(fontSource)
	if err != nil {
		return err
	}

	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)
	if err != nil {
		return err
	}

	//创建一个新的 Context，会对其设置一些默认值
	fc := freetype.NewContext()
	fc.SetDPI(72)              //设置屏幕每英寸的分辨率
	fc.SetFont(trueTypeFont)   //设置用于绘制文本的字体
	fc.SetFontSize(d.Size0)    //以磅为单位设置字体大小
	fc.SetClip(d.JPG.Bounds()) //设置剪裁矩形以进行绘制
	fc.SetDst(d.JPG)           //设置目标图像
	fc.SetSrc(image.Black)     //设置绘制操作的源图像，通常为 image.Uniform

	pt := freetype.Pt(d.X0, d.Y0)
	_, err = fc.DrawString(d.Title, pt)
	if err != nil {
		return err
	}

	fc.SetFontSize(d.Size1)
	_, err = fc.DrawString(d.SubTitle, freetype.Pt(d.X1, d.Y1)) //根据 Pt 的坐标值绘制给定的文本内容
	if err != nil {
		return err
	}

	err = jpeg.Encode(d.Merged, d.JPG, nil)
	if err != nil {
		return err
	}
	return nil
}
