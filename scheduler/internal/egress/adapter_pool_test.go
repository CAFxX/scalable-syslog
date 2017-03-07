package egress_test

import (
	"net"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	v1 "github.com/cloudfoundry-incubator/scalable-syslog/api/v1"
	"github.com/cloudfoundry-incubator/scalable-syslog/scheduler/internal/egress"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AdapterPool", func() {
	It("returns a list of adapter clients", func() {
		addr, cleanup := startGRPCServer()
		defer cleanup()

		pool := egress.NewAdapterPool([]string{addr}, grpc.WithInsecure())
		client := pool[0]

		_, err := client.ListBindings(context.Background(), &v1.ListBindingsRequest{})
		Expect(err).NotTo(HaveOccurred())
	})

	It("returns a pool with a unconnected client", func() {
		pool := egress.NewAdapterPool([]string{"0.0.0.0:1234"}, grpc.WithInsecure())
		client := pool[0]

		_, err := client.ListBindings(context.Background(), &v1.ListBindingsRequest{})
		Expect(err).To(HaveOccurred())
	})

	Context("Sub", func() {
		addr := "1.1.1.1"

		It("returns subset of the pool", func() {
			pool := egress.NewAdapterPool([]string{addr, addr, addr, addr}, grpc.WithInsecure())
			subset := pool.Sub(1, 2)

			Expect(len(subset)).To(Equal(2))
		})

		It("wraps back to the first index on overflow", func() {
			pool := egress.NewAdapterPool([]string{addr, addr, addr}, grpc.WithInsecure())
			subset := pool.Sub(1, 3)

			Expect(len(subset)).To(Equal(3))
		})

		It("returns all clients when more are requested than available", func() {
			pool := egress.NewAdapterPool([]string{addr, addr}, grpc.WithInsecure())
			subset := pool.Sub(0, 3)

			Expect(len(subset)).To(Equal(2))
		})
	})
})

func startGRPCServer() (string, func()) {
	lis, err := net.Listen("tcp", "localhost:0")
	Expect(err).NotTo(HaveOccurred())
	testServer := NewTestAdapterServer()
	grpcServer := grpc.NewServer()
	v1.RegisterAdapterServer(grpcServer, testServer)

	go grpcServer.Serve(lis)

	return lis.Addr().String(), func() {
		grpcServer.Stop()
		lis.Close()
	}
}