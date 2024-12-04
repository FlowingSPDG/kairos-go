package kairos

import (
	"fmt"
	"net"
)

// TODO: Support https

type Endpoints struct {
	protocol string
	ip       string
	port     string
}

func NewEndpoints(ip, port string) *Endpoints {
	return &Endpoints{
		protocol: "http",
		ip:       ip,
		port:     port,
	}
}

type AuxIdentifier interface {
	~int | ~string
}

func (e *Endpoints) Auxs() string {
	return fmt.Sprintf("%s://%s/aux", e.protocol, net.JoinHostPort(e.ip, e.port))
}

func endPointAux[T AuxIdentifier](e *Endpoints, id T) string {
	return fmt.Sprintf("%s://%s/aux/%v", e.protocol, net.JoinHostPort(e.ip, e.port), id)
}

func (e *Endpoints) AuxByID(id string) string {
	return endPointAux(e, id)
}
func (e *Endpoints) AuxByNumber(number int) string {
	return endPointAux(e, number)
}

func (e *Endpoints) Inputs() string {
	return fmt.Sprintf("%s://%s/inputs", e.protocol, net.JoinHostPort(e.ip, e.port))
}

func endPointInput[T InputIdentifier](e *Endpoints, id T) string {
	return fmt.Sprintf("%s://%s/inputs/%v", e.protocol, net.JoinHostPort(e.ip, e.port), id)
}

func (e *Endpoints) InputByID(id string) string {
	return endPointInput(e, id)
}
func (e *Endpoints) InputByNumber(number int) string {
	return endPointInput(e, number)
}

func (e *Endpoints) Macros() string {
	return fmt.Sprintf("%s://%s/macros", e.protocol, net.JoinHostPort(e.ip, e.port))
}

func (e *Endpoints) Macro(id string) string {
	return fmt.Sprintf("%s://%s/macros/%s", e.protocol, net.JoinHostPort(e.ip, e.port), id)
}

func (e *Endpoints) Multiviewers() string {
	return fmt.Sprintf("%s://%s/multiviewers", e.protocol, net.JoinHostPort(e.ip, e.port))
}

func endPointMultiviewers[T MultiviewerIdentifier](e *Endpoints, id T) string {
	return fmt.Sprintf("%s://%s/multiviewers/%v", e.protocol, net.JoinHostPort(e.ip, e.port), id)
}

func (e *Endpoints) MultiviewerByID(id string) string {
	return endPointMultiviewers(e, id)
}
func (e *Endpoints) MultiviewerInputByNumber(number int) string {
	return endPointMultiviewers(e, number)
}

func (e *Endpoints) Scenes() string {
	return fmt.Sprintf("%s://%s/scenes", e.protocol, net.JoinHostPort(e.ip, e.port))
}

func (e *Endpoints) Scene(id string) string {
	return fmt.Sprintf("%s://%s/scenes/%s", e.protocol, net.JoinHostPort(e.ip, e.port), id)
}
