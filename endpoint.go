package kairos

import (
	"fmt"
	"net"
)

// TODO: Support https

type Endpoints struct {
	ip   string
	port string
}

func NewEndpoints(ip, port string) *Endpoints {
	return &Endpoints{
		ip:   ip,
		port: port,
	}
}

func (e *Endpoints) Inputs() string {
	return fmt.Sprintf("http://%s/inputs", net.JoinHostPort(e.ip, e.port))
}

func endPointInput[T InputIdentifier](e *Endpoints, id T) string {
	return fmt.Sprintf("http://%s/inputs/%v", net.JoinHostPort(e.ip, e.port), id)
}

func (e *Endpoints) InputByID(id string) string {
	return endPointInput(e, id)
}
func (e *Endpoints) InputByNumber(number int) string {
	return endPointInput(e, number)
}

func (e *Endpoints) Macros() string {
	return fmt.Sprintf("http://%s/macros", net.JoinHostPort(e.ip, e.port))
}

func (e *Endpoints) Macro(id string) string {
	return fmt.Sprintf("http://%s/macros/%s", net.JoinHostPort(e.ip, e.port), id)
}

func (e *Endpoints) Multiviewers() string {
	return fmt.Sprintf("http://%s/multiviewers", net.JoinHostPort(e.ip, e.port))
}

func endPointMultiviewers[T MultiviewerIdentifier](e *Endpoints, id T) string {
	return fmt.Sprintf("http://%s/multiviewers/%v", net.JoinHostPort(e.ip, e.port), id)
}

func (e *Endpoints) MultiviewerByID(id string) string {
	return endPointMultiviewers(e, id)
}
func (e *Endpoints) MultiviewerInputByNumber(number int) string {
	return endPointMultiviewers(e, number)
}

func (e *Endpoints) Scenes() string {
	return fmt.Sprintf("http://%s/scenes", net.JoinHostPort(e.ip, e.port))
}

func (e *Endpoints) Scene(id string) string {
	return fmt.Sprintf("http://%s/scenes/%s", net.JoinHostPort(e.ip, e.port), id)
}