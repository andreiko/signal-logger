package main

import (
	"os"
	"github.com/Sirupsen/logrus"
	"os/signal"
	"syscall"
)

var logger = logrus.New()
var log *logrus.Entry = logger.WithField("version", Version)

func main() {
	cloudwatchGroup := os.Getenv("CLOUDWATCH_GROUP")
	cloudwatchStream := os.Getenv("CLOUDWATCH_STREAM")

	if cloudwatchGroup == "" {
		log.Warn("Skipping InstallCloudwatch as CLOUDWATCH_GROUP is not set")
	} else {
		if cloudwatchStream == "" {
			if hostname, ok := os.LookupEnv("HOSTNAME"); ok && hostname != "" {
				cloudwatchStream = hostname
			} else {
				log.Fatal("Clould not configure CloudWatch: neither CLOUDWATCH_STREAM nor HOSTNAME is set")
			}
		}

		if err := InstallCloudwatch(cloudwatchGroup, cloudwatchStream); err != nil {
			log.WithError(err).Fatal("Could not install CloudWatch hook")
		}
	}

	output := make(chan os.Signal, 1)

	signals := []os.Signal{
		syscall.SIGABRT,
		syscall.SIGALRM,
		syscall.SIGBUS,
		syscall.SIGCHLD,
		syscall.SIGCONT,
		syscall.SIGEMT,
		syscall.SIGFPE,
		syscall.SIGHUP,
		syscall.SIGILL,
		syscall.SIGINFO,
		syscall.SIGINT,
		syscall.SIGIO,
		syscall.SIGIOT,
		syscall.SIGKILL,
		syscall.SIGPIPE,
		syscall.SIGPROF,
		syscall.SIGQUIT,
		syscall.SIGSEGV,
		syscall.SIGSTOP,
		syscall.SIGSYS,
		syscall.SIGTERM,
		syscall.SIGTRAP,
		syscall.SIGTSTP,
		syscall.SIGTTIN,
		syscall.SIGTTOU,
		syscall.SIGURG,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		syscall.SIGVTALRM,
		syscall.SIGWINCH,
		syscall.SIGXCPU,
		syscall.SIGXFSZ,
	}
	for _, s := range signals {
		signal.Notify(output, s)
	}

	log.WithField("process_id", os.Getpid()).Info("started")
	for sig := range output {
		log.WithField("code", int(sig.(syscall.Signal))).Info(sig.String())
	}
}
