package avs

import (
	"strings"
	"time"
)

// AlertType specifies the type of an alert.
type AlertType string

// Possible values for AlertType, used by the SetAlert directive.
const (
	// AlertTypeAlarm specifies the type for an alarm. Alarms are scheduled for
	// specific times (e.g., to wake the user up).
	AlertTypeAlarm = AlertType("ALARM")
	// AlertTypeTimer specifies the type for a timer. Timers count down a certain
	// amount of time (e.g., "timer for 5 minutes").
	AlertTypeTimer = AlertType("TIMER")
)

// AudioItem represents an attached or streamable audio item.
type AudioItem struct {
	AudioItemId string `json:"audioItemId"`
	Stream      Stream `json:"stream"`
}

// ClearBehavior specifies how the play queue should be cleared.
type ClearBehavior string

// Possible values for ClearBehavior.
const (
	// ClearBehaviorClearAll clears all items in the queue, including the current
	// one.
	ClearBehaviorClearAll = ClearBehavior("CLEAR_ALL")
	// ClearBehaviorClearEnqueued clears all queued items in the queue (not
	// including the current one).
	ClearBehaviorClearEnqueued = ClearBehavior("CLEAR_ENQUEUED")
)

// PlayBehavior specifies how an audio item should be inserted into the play
// queue.
type PlayBehavior string

// Possible values for PlayBehavior.
const (
	// PlayBehaviorEnqueue specifies that the audio should be played after current
	// queue of audio.
	PlayBehaviorEnqueue = PlayBehavior("ENQUEUE")
	// PlayBehaviorReplaceAll specifies that the audio should play immediately,
	// throwing away all queued audio.
	PlayBehaviorReplaceAll = PlayBehavior("REPLACE_ALL")
	// PlayBehaviorReplaceEnqueued specifies that the audio should play after the
	// currently playing audio, replacing all queued audio.
	PlayBehaviorReplaceEnqueued = PlayBehavior("REPLACE_ENQUEUED")
)

// PlayerActivity specifies what state the audio player is in.
type PlayerActivity string

// Possible values for PlayerActivity.
const (
	PlayerActivityBufferUnderrun = PlayerActivity("BUFFER_UNDERRUN")
	PlayerActivityIdle           = PlayerActivity("IDLE")
	PlayerActivityPaused         = PlayerActivity("PAUSED")
	PlayerActivityPlaying        = PlayerActivity("PLAYING")
	PlayerActivityFinished       = PlayerActivity("FINISHED")
)

type ProgressReport struct {
	ProgressReportIntervalInMilliseconds float64 `json:"progressReportIntervalInMilliseconds"`
}

func (p *ProgressReport) Interval() time.Duration {
	return time.Duration(p.ProgressReportIntervalInMilliseconds) * time.Millisecond
}

// An audio stream which can either be attached with the response or a remote URL.
type Stream struct {
	ExpiryTime           string         `json:"expiryTime"`
	OffsetInMilliseconds float64        `json:"offsetInMilliseconds"`
	ProgressReport       ProgressReport `json:"progressReport"`
	Token                string         `json:"token"`
	URL                  string         `json:"url"`
}

// ContentId returns the content id of the audio, if it's attached with the
// response; otherwise, an empty string.
func (s *Stream) ContentId() string {
	if !strings.HasPrefix(s.URL, "cid:") {
		return ""
	}
	return s.URL[4:]
}
