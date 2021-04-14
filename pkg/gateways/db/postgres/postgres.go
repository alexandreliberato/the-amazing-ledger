package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func ConnectPool(dbURL string, log *logrus.Logger) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}
	if log != nil {
		// use current directory to match stack frame
		// only for this package
		log.SetLevel(logrus.DebugLevel)
		//		log.SetFormatter(&SimpleFormater{pathMatch: wd})
		log.SetReportCaller(true)

		config.ConnConfig.Logger = logrusadapter.NewLogger(log)
	}
	db, err := pgxpool.ConnectConfig(context.Background(), config)
	return db, err
}

// SimpleFormater used only to demonstrate how to extract
// stack frame and get file name and function name
type SimpleFormater struct {
	pathMatch string
	logrus.TextFormatter
}

// Format log message and add extra fields to logger
func (f *SimpleFormater) Format(entry *logrus.Entry) ([]byte, error) {
	pc := make([]uintptr, 10)
	if n := runtime.Callers(0, pc); n != 0 {
		frames := runtime.CallersFrames(pc[:n])
		for {
			frame, more := frames.Next()
			if strings.Contains(frame.File, f.pathMatch) {
				entry.Data["file"] = frame.File + ":" + strconv.Itoa(frame.Line)
				entry.Data["func"] = frame.Func.Name()
			}
			if !more {
				break
			}
		}
	}

	serialized, err := json.Marshal(entry.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return append(serialized, '\n'), nil
}
