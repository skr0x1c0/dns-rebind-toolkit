package main

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func init() {
	logger, err := zap.NewDevelopment()
	AssertOk(err)
	Logger = logger.Sugar()
}
