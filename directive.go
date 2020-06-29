package avs

import (
	"strings"
	"time"
)

/********** Alerts **********/

// The DeleteAlert directive.
type DeleteAlert struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

// The SetAlert directive.
type SetAlert struct {
	*Message
	Payload Alert `json:"payload"`
}

/********** AudioPlayer **********/

// The ClearQueue directive.
type ClearQueue struct {
	*Message
	Payload struct {
		ClearBehavior ClearBehavior `json:"clearBehavior"`
	} `json:"payload"`
}

// The Play directive.
type Play struct {
	*Message
	Payload struct {
		AudioItem    AudioItem    `json:"audioItem"`
		PlayBehavior PlayBehavior `json:"playBehavior"`
	} `json:"payload"`
}

// The Stop directive.
type Stop struct {
	*Message
	Payload struct{} `json:"payload"`
}

/********** Speaker **********/

// The AdjustVolume directive.
type AdjustVolume struct {
	*Message
	Payload struct {
		Volume int `json:"volume"`
	} `json:"payload"`
}

// The SetMute directive.
type SetMute struct {
	*Message
	Payload struct {
		Mute bool `json:"mute"`
	} `json:"payload"`
}

// The SetVolume directive.
type SetVolume struct {
	*Message
	Payload struct {
		Volume int `json:"volume"`
	} `json:"payload"`
}

/********** SpeechRecognizer **********/

// The ExpectSpeech directive.
type ExpectSpeech struct {
	*Message
	Payload struct {
		TimeoutInMilliseconds int `json:"timeoutInMilliseconds"`
	} `json:"payload"`
}

func (m *ExpectSpeech) Timeout() time.Duration {
	return time.Duration(m.Payload.TimeoutInMilliseconds) * time.Millisecond
}

// The StopCapture directive.
type StopCapture struct {
	*Message
	Payload struct{} `json:"payload"`
}

/********** SpeechSynthesizer **********/

// The Speak directive.
type Speak struct {
	*Message
	Payload struct {
		Format string `json:"format"`
		URL    string `json:"url"`
		Token  string `json:"token"`
	} `json:"payload"`
}

func (m *Speak) ContentId() string {
	if !strings.HasPrefix(m.Payload.URL, "cid:") {
		return ""
	}
	return m.Payload.URL[4:]
}

/********** System **********/

// The SetEndpoint directive.
type SetEndpoint struct {
	*Message
	Payload struct {
		Endpoint string `json:"endpoint"`
	} `json:"payload"`
}

// The ResetUserInactivity directive.
type ResetUserInactivity struct {
	*Message
	Payload struct{} `json:"payload"`
}

type RenderTemplateTitle struct {
    MainTitle string `json:"mainTitle,omitempty"`
    SubTitle string `json:"subTitle,omitempty"`
}

type TempValues struct {
	Value string `json:"value,omitempty"`
	Arrow Image `json:"arrow,omitempty"`
}

type ListItem struct {
	LeftTextField string `json:"leftTextField,omitempty"`
    RightTextField string `json:"rightTextField,omitempty"`
}

type ImageDef struct {
	Url string `json:"url,omitempty"`
	DarkBackgroundUrl string `json:"darkBackgroundUrl,omitempty"`
	Size string `json:"size,omitempty"`
	WidthPixels int64 `json:"widthPixels,omitempty"`
	HeightPixels int64 `json:"heightPixels,omitempty"`
}

type Image struct {
  ContentDescription string `json:"contentDescription,omitempty"`
  Sources []ImageDef `json:"sources,omitempty"`
}

type Header struct{
	  Token string `json:"token"`
      Type string `json:"type"`

}

type RenderHeader struct {
	  Namespace string `json:"namespace"`
      Name string  `json:"name"`
      MessageId string `json:"messageId"`
      DialogRequestId string `json:"dialogRequestId"`
  }

type TypedTemplate interface {
	GetTemplate()
}

type Template struct {
	 RenderHeader
	 Payload struct {
	 RenderTemplateTitle `json:"title,omitempty"`
	} `json:"payload,omitempty"`
}

// GetMessage returns a pointer to the underlying Message object.
func (t *Template) GetTemplate() *Template {
	return t
}
type Bodytemplate1 struct {
	*Template
	SkillIcon  Image `json:"skillicon,omitempty"`
	TextField string `json:"textfield,omitempty"`
}

type Bodytemplate2 struct {
	*Template
	SkillIcon  Image `json:"skillicon,omitempty"`
	TextField string `json:"textfield,omitempty"`
	Image Image `json:"image,omitempty"`
}

type ListTemplate1 struct {
	*Template
    SkillIcon  Image `json:"skillicon,omitempty"`
    ListItems []ListItem  `json:"listItems,omitempty"`
}

type WeatherForecastDay struct {
	 Image Image `json:"image,omitempty"`
     Day string `json:"day,omitempty"`
     Date string `json:"date,omitempty"`
     HighTemperature string `json:"highTemperature,omitempty"`
     LowTemperature string `json:"lowTemperature,omitempty"`
}

type WeatherTemplate struct {
	Template
	SkillIcon  Image `json:"skillicon,omitempty"`
	CurrentWeather string `json:"currentWeather,omitempty"`
	Description string `json:"description,omitempty"`
	CurrentWeatherIcon Image `json:"currentWeatherIcon,omitempty"`
	HighTemperature TempValues `json:"highTemperature,omitempty"`
	LowTemperature TempValues `json:"lowTemperature,omitempty"`
	WeatherForecast []WeatherForecastDay `json:"weatherForecast,omitempty"`
}

// The ResetUserInactivity directive.
type RenderTemplate struct {
	*Message
	Payload struct {
	Token string `json:"token"`
    Type string `json:"type"`
  //  Template
    RenderTemplateTitle `json:"title,omitempty"`
    SkillIcon  Image `json:"skillicon,omitempty"`
	TextField string `json:"textfield,omitempty"`
	Image Image `json:"image,omitempty"`
	ListItems []ListItem  `json:"listItems,omitempty"`
	CurrentWeather string `json:"currentWeather,omitempty"`
	Description string `json:"description,omitempty"`
	CurrentWeatherIcon Image `json:"currentWeatherIcon,omitempty"`
	HighTemperature TempValues `json:"highTemperature,omitempty"`
	LowTemperature TempValues `json:"lowTemperature,omitempty"`
	WeatherForecast []WeatherForecastDay `json:"weatherForecast,omitempty"`
	} `json:"payload"`
}
/*func (m *RenderTemplate) ContentId() string {
	return m.Payload
}*/