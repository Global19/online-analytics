package useranalytics

import (
	"strings"
	"testing"

	"github.com/openshift/origin/pkg/client/testclient"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/kubernetes/pkg/api"
	ktestclient "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/fake"
)

func TestWoopraDestination(t *testing.T) {
	pod := &api.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "bar",
		},
	}

	// TODO:  this needs some kind of factory per object
	// that creates analyticsEvent objects
	event, _ := newEvent(api.Scheme, pod, watch.Added)

	dest := &WoopraDestination{
		Method:   "GET",
		Endpoint: "http://www.woopra.com/track/ce",
		Domain:   "dev.example.com",
		Client:   &mockHttpClient{},
	}

	err := dest.Send(event)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	mockClient := dest.Client.(*mockHttpClient)

	if !strings.Contains(mockClient.url, dest.Endpoint) {
		t.Errorf("Expected the destination to include %s , but got: %s", dest.Endpoint, mockClient.url)
	}

}

func TestWoopraLive(t *testing.T) {
	oc := &testclient.Fake{}
	kc := &ktestclient.Clientset{}

	items := WatchFuncList(kc, oc)

	for _, w := range items {
		m, err := meta.Accessor(w.objType)
		if err != nil {
			t.Errorf("Unable to create object meta for %v", w.objType)
		}
		m.SetName("foo")
		m.SetNamespace("foobar")

		event, _ := newEvent(api.Scheme, w.objType, watch.Added)

		dest := &WoopraDestination{
			Method:   "GET",
			Endpoint: "http://www.woopra.com/track/ce",
			Domain:   "dev.example.com",
			Client:   NewSimpleHttpClient(),
		}

		err = dest.Send(event)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestPrepEndpoint(t *testing.T) {
	before := "http://www.woopra.com/track/ce"
	expected := "http://www.woopra.com/track/ce?%s"
	after := prepEndpoint(before)
	if after != expected {
		t.Errorf("Expected %s, but got %s", expected, after)
	}
}
