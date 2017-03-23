package prepare

import (
	"bytes"
	"github.com/skatteetaten/architect/pkg/java/config"
	"testing"
)

func TestWriteStartscript(t *testing.T) {

	const mainClass string = "foo.bar.Main"
	const jvmOpts string = "-Dfoo=bar"
	const applicationArgs string = "--logging.config=logback.xml"

	classpath := []string{"/app/lib/metrics.jar", "/app/lib/rt.jar", "/app/lib/spring.jar"}

	cfg := &config.DeliverableMetadata{
		Docker: &struct {
			Maintainer string            `json:"maintainer"`
			Labels     map[string]string `json:"labels"`
		}{},
		Java: &struct {
			MainClass       string `json:"mainClass"`
			JvmOpts         string `json:"jvmOpts"`
			ApplicationArgs string `json:"applicationArgs"`
			ReadinessURL    string `json:"readinessUrl"`
		}{
			MainClass:       mainClass,
			JvmOpts:         jvmOpts,
			ApplicationArgs: applicationArgs,
		},
		Openshift: &struct {
			ReadinessURL              string `json:"readinessUrl"`
			ReadinessOnManagementPort string `json:"readinessOnManagementPort"`
		}{},
	}

	var buf bytes.Buffer

	NewJavaStartScript(classpath, cfg).Write(&buf)

	startscript := buf.String()

	assertContainsElement(t, startscript, mainClass)
	assertContainsElement(t, startscript, jvmOpts)
	assertContainsElement(t, startscript, applicationArgs)
	assertContainsElement(t, startscript, classpath[0])
	assertContainsElement(t, startscript, classpath[1])
	assertContainsElement(t, startscript, classpath[2])
}
