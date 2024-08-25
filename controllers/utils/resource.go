package utils

import (
	"bytes"
	"fmt"
	"text/template"

	v1 "kubebuilder-demo/api/v1"

	goyaml "gopkg.in/yaml.v2"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func parseTemplate(templateName string, app *v1.Myapp) []byte {
	tmpl, err := template.ParseFiles("../controllers/template/" + templateName + ".yml")
	if err != nil {
		panic(err)
	}
	b := new(bytes.Buffer)
	err = tmpl.Execute(b, app)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

func NewDeployment(app *v1.Myapp) *appv1.Deployment {
	d := &appv1.Deployment{}
	err := yaml.Unmarshal(parseTemplate("deployment", app), d)
	if err != nil {
		panic(err)
	}

	// 将 Deployment 结构体转换为 YAML 格式
	yamlData, err := goyaml.Marshal(d)
	if err != nil {
		panic(err)
	}

	// 输出 YAML 格式的 Deployment
	fmt.Printf("%s\n", yamlData)
	return d
}

func NewIngress(app *v1.Myapp) *netv1.Ingress {
	i := &netv1.Ingress{}
	err := yaml.Unmarshal(parseTemplate("ingress", app), i)
	if err != nil {
		panic(err)
	}
	return i
}

func NewService(app *v1.Myapp) *corev1.Service {
	s := &corev1.Service{}
	err := yaml.Unmarshal(parseTemplate("service", app), s)
	if err != nil {
		panic(err)
	}
	return s
}
