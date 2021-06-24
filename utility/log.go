package utility

import (
	"bytes"
	"context"
	"errors"
	"github.com/kpango/glg"
	_ "github.com/lestrrat-go/apache-logformat"
	"github.com/lestrrat-go/file-rotatelogs"
	"io"
	"log"
	"sync"
	"time"
)

type LEVEL uint8

var CDRCount int

type RotateWriter struct {
	writer io.Writer
	dur    time.Duration
	once   sync.Once
	cancel context.CancelFunc
	mu     sync.Mutex
	buf    *bytes.Buffer
}

func NewRotateWriter(w io.Writer, dur time.Duration, buf *bytes.Buffer) io.WriteCloser {
	return &RotateWriter{
		writer: w,
		dur:    dur,
		buf:    buf,
	}
}

func (r *RotateWriter) Write(b []byte) (int, error) {
	if r.buf == nil || r.writer == nil {
		return 0, errors.New("error invalid rotate config")
	}
	r.once.Do(func() {
		var ctx context.Context
		ctx, r.cancel = context.WithCancel(context.Background())
		go func() {
			tick := time.NewTicker(r.dur)
			for {
				select {
				case <-ctx.Done():
					tick.Stop()
					return
				case <-tick.C:
					r.mu.Lock()
					r.writer.Write(r.buf.Bytes())
					r.buf.Reset()
					r.mu.Unlock()
				}
			}
		}()
	})
	r.mu.Lock()
	r.buf.Write(b)
	r.mu.Unlock()
	return len(b), nil
}

func (r *RotateWriter) Close() error {
	if r.cancel != nil {
		r.cancel()
	}
	return nil
}

func Initialize() {
	infolog := glg.FileWriter("info.log", 0666)
	CDR := "CDR"
	errlog := glg.FileWriter("error.log", 0666)
	warnlog := glg.FileWriter("warn.log", 0666)
	//rotate := NewRotateWriter(os.Stdin, time.Second*10, bytes.NewBuffer(make([]byte, 0, 4096)))

	//defer infolog.Close()
	//defer errlog.Close()
	//defer rotate.Close()

	glg.Get().
		SetMode(glg.BOTH). // default is STD
		// DisableColor().
		// SetMode(glg.NONE).
		// SetMode(glg.WRITER).
		// SetMode(glg.BOTH).
		// InitWriter().
		// AddWriter(customWriter).
		// SetWriter(customWriter).
		// AddLevelWriter(glg.LOG, customWriter).
		// AddLevelWriter(glg.INFO, customWriter).
		// AddLevelWriter(glg.WARN, customWriter).
		// AddLevelWriter(glg.ERR, customWriter).
		// SetLevelWriter(glg.LOG, customWriter).
		// SetLevelWriter(glg.INFO, customWriter).
		// SetLevelWriter(glg.WARN, customWriter).
		// SetLevelWriter(glg.ERR, customWriter).
		// EnableJSON().
		AddLevelWriter(glg.INFO, infolog).                   // add info log file destination
		AddLevelWriter(glg.ERR, errlog).                     // add error log file destination
		AddLevelWriter(glg.WARN, warnlog).                   // add error log file destination
		AddStdLevel(CDR, glg.STD, true).                     //user custom log level
		SetLevelColor(glg.TagStringToLevel(CDR), glg.Orange) // set color output to user custom level

	rl, _ := rotatelogs.New(
		"CDR/cdr.%Y%m%d%H%M",
		rotatelogs.WithLinkName("CDR/cdr"),
		rotatelogs.WithMaxAge(24*time.Hour),
		// rotatelogs.WithRotationTime(time.Hour),
		// rotatelogs.WithRotationSize(1024*1024),
		rotatelogs.ForceNewFile(),
	)

	log.SetOutput(rl)

	/* elsewhere ... */
	glg.Debug("Log Initialized successfully")
}
func Log(level string, val ...interface{}) {
	switch level {
	case "DEBUG":
		glg.Debug(val)
	case "WARNING":
		glg.Warn(val)
	case "INFO":
		glg.Info(val)
	case "ERROR":
		glg.Error(val)
	default:
		glg.Debug(val)
	}
}

func LogCDR(cdr CDR) {
	log.Printf(cdr.Log())
}
