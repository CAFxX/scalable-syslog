package metric_test

import (
	"log"
	"net"
	"testing"

	"code.cloudfoundry.org/scalable-syslog/internal/api/loggregator/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMetric(t *testing.T) {
	grpclog.SetLogger(log.New(GinkgoWriter, "", 0))
	log.SetOutput(GinkgoWriter)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Metric Suite")
}

func startIngressServer() (string, *SpyIngressServer) {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	spyIngressServer := &SpyIngressServer{
		senders: make(chan loggregator_v2.Ingress_SenderServer, 100),
	}

	s := grpc.NewServer()
	loggregator_v2.RegisterIngressServer(s, spyIngressServer)
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	return lis.Addr().String(), spyIngressServer
}

func rxToCh(rx loggregator_v2.Ingress_SenderServer) <-chan *loggregator_v2.Envelope {
	c := make(chan *loggregator_v2.Envelope, 100)
	go func() {
		for {
			e, err := rx.Recv()
			if err != nil {
				return
			}
			c <- e
		}
	}()
	return c
}

type SpyIngressServer struct {
	senders chan loggregator_v2.Ingress_SenderServer
}

func (s *SpyIngressServer) Sender(sender loggregator_v2.Ingress_SenderServer) error {
	s.senders <- sender
	<-sender.Context().Done()
	return sender.Context().Err()
}

func fetchReceiver(spyIngressServer *SpyIngressServer) (rx loggregator_v2.Ingress_SenderServer) {
	Eventually(spyIngressServer.senders, 3).Should(Receive(&rx))
	return rx
}
