package client

import (
	"k8s.io/client-go/rest"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
)

func assert(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
func check(t *testing.T, err error) {
	if err != nil {
		t.Errorf("get error: %+v", err)
	}
}

func getConfig() *rest.Config {
	return &rest.Config{
		Host:        "https://119.28.224.65:6443",
		BearerToken: "fuckfk.baidu12345678910",
		APIPath:     "apis",
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
}

func TestGetClient(t *testing.T) {
	c, err := NewKubeClient(getConfig(), "dev")
	if err != nil {
		t.Errorf("get client error: %s", err.Error())
	}

	version, err := c.GetVersionByGroup("apps")
	check(t, err)
	assert(t, version, "v1")

}

func TestList(t *testing.T) {
	c, _ := NewKubeClient(getConfig(), "dev")

	result, err := c.ListResource("default", metav1.ListOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		}})
	check(t, err)
	fmt.Println(fmt.Sprintf("result: %+v %d", result.Object, len(result.Items)))
}

func TestGet(t *testing.T) {
	c, _ := NewKubeClient(getConfig(), "dev")
	result, err := c.GetResource("default", "kubernetes", metav1.GetOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
	})
	check(t, err)
	fmt.Println(result)
}
