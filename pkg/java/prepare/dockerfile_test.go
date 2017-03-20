package prepare

import (
	"os"
	"bytes"
	"testing"
	"strings"
	"github.com/Skatteetaten/architect/pkg/java/config"
)

func TestBuild(t *testing.T) {

	const maintainer string = "tester@skatteetaten.no"
	const k8sDescription string = "Demo application with spring boot on Openshift."
	const openshiftTags string = "openshift,springboot"
	const readinessUrl string = "http://ready.skead.no"
	const envVar01 string = "1.0.95-b2.2.3-oracle8-1.4.0"
	const envVar02 string = "1.0.95"

	cfg := &config.ArchitectConfig{
		Docker: &struct {
			Maintainer string `json:"maintainer"`
			Labels     map[string]string `json:"labels"`
		}{
			Maintainer: maintainer,
			Labels: map[string]string {
				"io.k8s.description": k8sDescription,
				"io.openshift.tags":  openshiftTags,
			},
		},
		Openshift: &struct {
			ReadinessURL              string `json:"readinessUrl"`
			ReadinessOnManagementPort string `json:"readinessOnManagementPort"`
		}{
			ReadinessURL: readinessUrl,
		},
	}

	var buf bytes.Buffer
	Env := make(map[string]string)
	Env["ENV_VAR_01"] = envVar01
	Env["ENV_VAR_02"] = envVar02


	NewForConfig("BaseUrl", Env, cfg).Build(&buf)
	NewForConfig("BaseUrl", Env, cfg).Build(os.Stdout)

	dockerfile := buf.String()

	assertContainsElement(t, dockerfile, maintainer)
	assertContainsElement(t, dockerfile, k8sDescription)
	assertContainsElement(t, dockerfile, openshiftTags)
	assertContainsElement(t, dockerfile, readinessUrl)
	assertContainsElement(t, dockerfile, envVar01)
	assertContainsElement(t, dockerfile, envVar02)
}

func assertContainsElement(t *testing.T, target string, element string) {
	if strings.Contains(target, element) == false {
		t.Error( "excpected", element, ", got", target)
	}
}
