package avs

import (
	"time"
)
const AlertFamily = "Alerts"
const AlertEnteredBackgroundEvent = "AlertEnteredBackground"
const AlertEnteredForegroundEvent = "AlertEnteredForeground"
const AlertStartedEvent = "AlertStarted"
const AlertStoppedEvent = "AlertStopped"
const DeleteAlertFailedEvent = "DeleteAlertFailed"
const DeleteAlertSucceededEvent = "DeleteAlertSucceeded"
const SetAlertFailedEvent = "SetAlertFailed"
const SetAlertSucceededEvent = "SetAlertSucceeded"
const AudioPlayerFamily = "AudioPlayer"
const PlaybackFailedEvent = "PlaybackFailed"
const PlaybackFinishedEvent = "PlaybackFinished"
const PlaybackNearlyFinishedEvent = "PlaybackNearlyFinished"
const PlaybackPausedEvent = "PlaybackPaused"
const PlaybackQueueClearedEvent = "PlaybackQueueCleared"
const PlaybackResumedEvent = "PlaybackResumed"
const PlaybackStartedEvent = "PlaybackStarted"
const PlaybackStoppedEvent = "PlaybackStopped"
const PlaybackStutterStartedEvent = "PlaybackStutterStarted"
const PlaybackStutterFinishedEvent = "PlaybackStutterFinished"
const ProgressReportDelayElapsedEvent = "ProgressReportDelayElapsed"
const ProgressReportIntervalElapsedEvent = "ProgressReportIntervalElapsed"
const StreamMetadataExtractedEvent = "StreamMetadataExtracted"
const PlaybackControllerFamily = "PlaybackController"
const NextCommandIssuedEvent = "NextCommandIssued"
const PauseCommandIssuedEvent = "PauseCommandIssued"
const PlayCommandIssuedEvent = "PlayCommandIssued"
const PreviousCommandIssuedEvent = "PreviousCommandIssued"
const SpeakerFamily = "Speaker"
const MuteChangedEvent = "MuteChanged"
const VolumeChangedEvent = "VolumeChanged"
const SpeechRecognizerFamily = "SpeechRecognizer"
const ExpectSpeechTimedOutEvent = "ExpectSpeechTimedOut"
const RecognizeEvent = "Recognize"
const SpeechSynthesizerFamily = "SpeechSynthesizer"
const SpeechFinishedEvent = "SpeechFinished"
const SpeechStartedEvent = "SpeechStarted"
const SettingsFamily = "Settings"
const SettingsUpdatedEvent = "SettingsUpdated"
const SystemFamily = "System"
const ExceptionEncounteredEvent = "ExceptionEncountered"
const SynchronizeStateEvent = "SynchronizeState"
const UserInactivityReportEvent = "UserInactivityReport"


// newEvent creates a Message suited for being used as an event value.
func newEvent(namespace, name, messageId, dialogRequestId string) *Message {
	m := &Message{
		Header: map[string]string{
			"namespace": namespace,
			"name":      name,
			"messageId": messageId,
		},
		Payload: nil,
	}
	if dialogRequestId != "" {
		m.Header["dialogRequestId"] = dialogRequestId
	}
	return m
}

/********** Alerts **********/

// The AlertEnteredBackground event.
type AlertEnteredBackground struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

func NewAlertEnteredBackground(messageId, token string) *AlertEnteredBackground {
	m := new(AlertEnteredBackground)
	m.Message = newEvent(AlertFamily, AlertEnteredBackgroundEvent, messageId, "")
	m.Payload.Token = token
	return m
}

// The AlertEnteredForeground event.
type AlertEnteredForeground struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

func NewAlertEnteredForeground(messageId, token string) *AlertEnteredForeground {
	m := new(AlertEnteredForeground)
	m.Message = newEvent(AlertFamily, AlertEnteredForegroundEvent, messageId, "")
	m.Payload.Token = token
	return m
}

// The AlertStarted event.
type AlertStarted struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

func NewAlertStarted(messageId, token string) *AlertStarted {
	m := new(AlertStarted)
	m.Message = newEvent(AlertFamily, AlertStartedEvent, messageId, "")
	m.Payload.Token = token
	return m
}

// The AlertStopped event.
type AlertStopped struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

func NewAlertStopped(messageId, token string) *AlertStopped {
	m := new(AlertStopped)
	m.Message = newEvent(AlertFamily, AlertStoppedEvent, messageId, "")
	m.Payload.Token = token
	return m
}

// The DeleteAlertFailed event.
type DeleteAlertFailed struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

func NewDeleteAlertFailed(messageId, token string) *DeleteAlertFailed {
	m := new(DeleteAlertFailed)
	m.Message = newEvent(AlertFamily, DeleteAlertFailedEvent, messageId, "")
	m.Payload.Token = token
	return m
}

// The DeleteAlertSucceeded event.
type DeleteAlertSucceeded struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

func NewDeleteAlertSucceeded(messageId, token string) *DeleteAlertSucceeded {
	m := new(DeleteAlertSucceeded)
	m.Message = newEvent(AlertFamily, DeleteAlertSucceededEvent, messageId, "")
	m.Payload.Token = token
	return m
}

// The SetAlertFailed event.
type SetAlertFailed struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

func NewSetAlertFailed(messageId, token string) *SetAlertFailed {
	m := new(SetAlertFailed)
	m.Message = newEvent(AlertFamily, SetAlertFailedEvent, messageId, "")
	m.Payload.Token = token
	return m
}

// The SetAlertSucceeded event.
type SetAlertSucceeded struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

func NewSetAlertSucceeded(messageId, token string) *SetAlertSucceeded {
	m := new(SetAlertSucceeded)
	m.Message = newEvent(AlertFamily, SetAlertSucceededEvent, messageId, "")
	m.Payload.Token = token
	return m
}

/********** AudioPlayer **********/
// Also used by the PlaybackState context.
type playbackState struct {
	Token                string         `json:"token"`
	OffsetInMilliseconds int            `json:"offsetInMilliseconds"`
	PlayerActivity       PlayerActivity `json:"playerActivity"`
}

// The PlaybackFailed event.
type PlaybackFailed struct {
	*Message
	Payload struct {
		Token                string        `json:"token"`
		CurrentPlaybackState playbackState `json:"currentPlaybackState"`
		Error struct {
			Type    MediaErrorType `json:"type"`
			Message string         `json:"message"`
		} `json:"error"`
	} `json:"payload"`
}

func NewPlaybackFailed(messageId, token string, errorType MediaErrorType, errorMessage string) *PlaybackFailed {
	m := new(PlaybackFailed)
	m.Message = newEvent(AudioPlayerFamily, PlaybackFailedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.Error.Type = errorType
	m.Payload.Error.Message = errorMessage
	return m
}

// The PlaybackFinished event.
type PlaybackFinished struct {
	*Message
	Payload struct {
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	} `json:"payload"`
}

func NewPlaybackFinished(messageId, token string, offset time.Duration) *PlaybackFinished {
	m := new(PlaybackFinished)
	m.Message = newEvent(AudioPlayerFamily, PlaybackFinishedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.OffsetInMilliseconds = int(offset.Seconds() * 1000)
	return m
}

// The PlaybackNearlyFinished event.
type PlaybackNearlyFinished struct {
	*Message
	Payload struct {
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	} `json:"payload"`
}

func NewPlaybackNearlyFinished(messageId, token string, offset time.Duration) *PlaybackNearlyFinished {
	m := new(PlaybackNearlyFinished)
	m.Message = newEvent(AudioPlayerFamily, PlaybackNearlyFinishedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.OffsetInMilliseconds = int(offset.Seconds() * 1000)
	return m
}

// The PlaybackPaused event.
type PlaybackPaused struct {
	*Message
	Payload struct {
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	} `json:"payload"`
}

func NewPlaybackPaused(messageId, token string, offset time.Duration) *PlaybackPaused {
	m := new(PlaybackPaused)
	m.Message = newEvent(AudioPlayerFamily, PlaybackPausedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.OffsetInMilliseconds = int(offset.Seconds() * 1000)
	return m
}

// The PlaybackQueueCleared event.
type PlaybackQueueCleared struct {
	*Message
	Payload struct{} `json:"payload"`
}

func NewPlaybackQueueCleared(messageId string) *PlaybackQueueCleared {
	m := new(PlaybackQueueCleared)
	m.Message = newEvent(AudioPlayerFamily, PlaybackQueueClearedEvent, messageId, "")
	return m
}

// The PlaybackResumed event.
type PlaybackResumed struct {
	*Message
	Payload struct {
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	} `json:"payload"`
}

func NewPlaybackResumed(messageId, token string, offset time.Duration) *PlaybackResumed {
	m := new(PlaybackResumed)
	m.Message = newEvent(AudioPlayerFamily, PlaybackResumedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.OffsetInMilliseconds = int(offset.Seconds() * 1000)
	return m
}

// The PlaybackStarted event.
type PlaybackStarted struct {
	*Message
	Payload struct {
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	} `json:"payload"`
}

func NewPlaybackStarted(messageId, token string, offset time.Duration) *PlaybackStarted {
	m := new(PlaybackStarted)
	m.Message = newEvent(AudioPlayerFamily, PlaybackStartedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.OffsetInMilliseconds = int(offset.Seconds() * 1000)
	return m
}

// The PlaybackStopped event.
type PlaybackStopped struct {
	*Message
	Payload struct {
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	} `json:"payload"`
}

func NewPlaybackStopped(messageId, token string, offset time.Duration) *PlaybackStopped {
	m := new(PlaybackStopped)
	m.Message = newEvent(AudioPlayerFamily, PlaybackStoppedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.OffsetInMilliseconds = int(offset.Seconds() * 1000)
	return m
}

// The PlaybackStutterStarted event.
type PlaybackStutterStarted struct {
	*Message
	Payload struct {
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	} `json:"payload"`
}

func NewPlaybackStutterStarted(messageId, token string, offset time.Duration) *PlaybackStutterStarted {
	m := new(PlaybackStutterStarted)
	m.Message = newEvent(AudioPlayerFamily, PlaybackStutterStartedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.OffsetInMilliseconds = int(offset.Seconds() * 1000)
	return m
}

// The PlaybackStutterFinished event.
type PlaybackStutterFinished struct {
	*Message
	Payload struct {
		Token                         string `json:"token"`
		OffsetInMilliseconds          int    `json:"offsetInMilliseconds"`
		StutterDurationInMilliseconds int    `json:"stutterDurationInMilliseconds"`
	} `json:"payload"`
}

func NewPlaybackStutterFinished(messageId, token string, offset, stutterDuration time.Duration) *PlaybackStutterFinished {
	m := new(PlaybackStutterFinished)
	m.Message = newEvent(AudioPlayerFamily, PlaybackStutterFinishedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.OffsetInMilliseconds = int(offset.Seconds() * 1000)
	m.Payload.StutterDurationInMilliseconds = int(stutterDuration.Seconds() * 1000)
	return m
}

// The ProgressReportDelayElapsed event.
type ProgressReportDelayElapsed struct {
	*Message
	Payload struct {
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	} `json:"payload"`
}

func NewProgressReportDelayElapsed(messageId, token string, offset time.Duration) *ProgressReportDelayElapsed {
	m := new(ProgressReportDelayElapsed)
	m.Message = newEvent(AudioPlayerFamily, ProgressReportDelayElapsedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.OffsetInMilliseconds = int(offset.Seconds() * 1000)
	return m
}

// The ProgressReportIntervalElapsed event.
type ProgressReportIntervalElapsed struct {
	*Message
	Payload struct {
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	} `json:"payload"`
}

func NewProgressReportIntervalElapsed(messageId, token string, offset time.Duration) *ProgressReportIntervalElapsed {
	m := new(ProgressReportIntervalElapsed)
	m.Message = newEvent(AudioPlayerFamily, ProgressReportIntervalElapsedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.OffsetInMilliseconds = int(offset.Seconds() * 1000)
	return m
}

// The StreamMetadataExtracted event.
type StreamMetadataExtracted struct {
	*Message
	Payload struct {
		Token    string                 `json:"token"`
		Metadata map[string]interface{} `json:"metadata"`
	} `json:"payload"`
}

func NewStreamMetadataExtracted(messageId, token string, metadata map[string]interface{}) *StreamMetadataExtracted {
	m := new(StreamMetadataExtracted)
	m.Message = newEvent(AudioPlayerFamily, StreamMetadataExtractedEvent, messageId, "")
	m.Payload.Token = token
	m.Payload.Metadata = metadata
	return m
}

/********** PlaybackController **********/

// The NextCommandIssued event.
type NextCommandIssued struct {
	*Message
	Payload struct{} `json:"payload"`
}

func NewNextCommandIssued(messageId string) *NextCommandIssued {
	m := new(NextCommandIssued)
	m.Message = newEvent(PlaybackControllerFamily, NextCommandIssuedEvent, messageId, "")
	return m
}

// The PauseCommandIssued event.
type PauseCommandIssued struct {
	*Message
	Payload struct{} `json:"payload"`
}

func NewPauseCommandIssued(messageId string) *PauseCommandIssued {
	m := new(PauseCommandIssued)
	m.Message = newEvent(PlaybackControllerFamily, PauseCommandIssuedEvent, messageId, "")
	return m
}

// The PlayCommandIssued event.
type PlayCommandIssued struct {
	*Message
	Payload struct{} `json:"payload"`
}

func NewPlayCommandIssued(messageId string) *PlayCommandIssued {
	m := new(PlayCommandIssued)
	m.Message = newEvent(PlaybackControllerFamily, PlayCommandIssuedEvent, messageId, "")
	return m
}

// The PreviousCommandIssued event.
type PreviousCommandIssued struct {
	*Message
	Payload struct{} `json:"payload"`
}

func NewPreviousCommandIssued(messageId string) *PreviousCommandIssued {
	m := new(PreviousCommandIssued)
	m.Message = newEvent(PlaybackControllerFamily, PreviousCommandIssuedEvent, messageId, "")
	return m
}

/********** Speaker **********/

// The MuteChanged event.
type MuteChanged struct {
	*Message
	Payload struct {
		Volume int  `json:"volume"`
		Muted  bool `json:"muted"`
	} `json:"payload"`
}

func NewMuteChanged(messageId string, volume int, muted bool) *MuteChanged {
	m := new(MuteChanged)
	m.Message = newEvent(SpeakerFamily, MuteChangedEvent, messageId, "")
	m.Payload.Volume = volume
	m.Payload.Muted = muted
	return m
}

// The VolumeChanged event.
type VolumeChanged struct {
	*Message
	Payload struct {
		Volume int  `json:"volume"`
		Muted  bool `json:"muted"`
	} `json:"payload"`
}

func NewVolumeChanged(messageId string, volume int, muted bool) *VolumeChanged {
	m := new(VolumeChanged)
	m.Message = newEvent(SpeakerFamily, VolumeChangedEvent, messageId, "")
	m.Payload.Volume = volume
	m.Payload.Muted = muted
	return m
}

/********** SpeechRecognizer **********/

// The ExpectSpeechTimedOut event.
type ExpectSpeechTimedOut struct {
	*Message
	Payload struct{} `json:"payload"`
}

func NewExpectSpeechTimedOut(messageId string) *ExpectSpeechTimedOut {
	m := new(ExpectSpeechTimedOut)
	m.Message = newEvent(SpeechRecognizerFamily, ExpectSpeechTimedOutEvent, messageId, "")
	return m
}


// RecognizeProfile identifies the ASR profile associated with your product.
type RecognizeProfile string

// Possible values for RecognizeProfile.
// Supports three distinct profiles optimized for speech at varying distances.
const (
	RecognizeProfileCloseTalk = RecognizeProfile("CLOSE_TALK")
	RecognizeProfileNearField = RecognizeProfile("NEAR_FIELD")
	RecognizeProfileFarField  = RecognizeProfile("FAR_FIELD")
)

// The Recognize event.
type Recognize struct {
	*Message
	Payload struct {
		Profile RecognizeProfile `json:"profile"`
		Format  string           `json:"format"`
	} `json:"payload"`
}

func NewRecognize(messageId, dialogRequestId string) *Recognize {
	return NewRecognizeWithProfile(messageId, dialogRequestId, RecognizeProfileCloseTalk)
}

func NewRecognizeWithProfile(messageId, dialogRequestId string, profile RecognizeProfile) *Recognize {
	m := new(Recognize)
	m.Message = newEvent(SpeechRecognizerFamily, RecognizeEvent, messageId, dialogRequestId)
	m.Payload.Format = "AUDIO_L16_RATE_16000_CHANNELS_1"
	m.Payload.Profile = profile
	return m
}

/********** SpeechSynthesizer **********/

// The SpeechFinished event.
type SpeechFinished struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

func NewSpeechFinished(messageId, token string) *SpeechFinished {
	m := new(SpeechFinished)
	m.Message = newEvent(SpeechSynthesizerFamily, SpeechFinishedEvent, messageId, "")
	m.Payload.Token = token
	return m
}

// The SpeechStarted event.
type SpeechStarted struct {
	*Message
	Payload struct {
		Token string `json:"token"`
	} `json:"payload"`
}

func NewSpeechStarted(messageId, token string) *SpeechStarted {
	m := new(SpeechStarted)
	m.Message = newEvent(SpeechSynthesizerFamily, SpeechStartedEvent, messageId, "")
	m.Payload.Token = token
	return m
}

/********** Settings **********/

// The SettingsUpdated event.
type Setting struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SettingsUpdated struct {
	*Message
	Payload struct {
		Settings []Setting `json:"settings"`
	} `json:"payload"`
}

type SettingLocale string
// Possible values for SettingLocale.
const (
	SettingLocaleUS = SettingLocale("en-US")
	SettingLocaleGB = SettingLocale("en-GB")
	SettingLocaleDE = SettingLocale("de-DE")
)

func NewLocaleSettingsUpdated(messageId string, locale SettingLocale) *SettingsUpdated {
	m := new(SettingsUpdated)
	m.Message = newEvent(SettingsFamily, SettingsUpdatedEvent, messageId, "")
	m.Payload.Settings = append(m.Payload.Settings, Setting{
		Key: "locale",
		Value: string(locale),
	})
	return m
}

/********** System **********/

// The ExceptionEncountered event.
type ExceptionEncountered struct {
	*Message
	Payload struct {
		UnparsedDirective string `json:"unparsedDirective"`
		Error             struct {
			Type    ErrorType `json:"type"`
			Message string    `json:"message"`
		} `json:"error"`
	} `json:"payload"`
}

func NewExceptionEncountered(messageId, directive string, errorType ErrorType, errorMessage string) *ExceptionEncountered {
	m := new(ExceptionEncountered)
	m.Message = newEvent(SystemFamily, ExceptionEncounteredEvent, messageId, "")
	m.Payload.UnparsedDirective = directive
	m.Payload.Error.Type = errorType
	m.Payload.Error.Message = errorMessage
	return m
}

// The SynchronizeState event.
type SynchronizeState struct {
	*Message
	Payload struct{} `json:"payload"`
}

func NewSynchronizeState(messageId string) *SynchronizeState {
	m := new(SynchronizeState)
	m.Message = newEvent(SystemFamily, SynchronizeStateEvent, messageId, "")
	return m
}

// The UserInactivityReport event.
type UserInactivityReport struct {
	*Message
	Payload struct {
		InactiveTimeInSeconds int `json:"inactiveTimeInSeconds"`
	} `json:"payload"`
}

func NewUserInactivityReport(messageId string, inactiveTime time.Duration) *UserInactivityReport {
	m := new(UserInactivityReport)
	m.Message = newEvent(SystemFamily, UserInactivityReportEvent, messageId, "")
	m.Payload.InactiveTimeInSeconds = int(inactiveTime.Seconds())
	return m
}
