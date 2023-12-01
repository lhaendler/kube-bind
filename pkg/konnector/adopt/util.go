package adopt

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

func InjectBoundAnnotation(obj *unstructured.Unstructured) {
	ann := obj.GetAnnotations()
	if ann == nil {
		ann = make(map[string]string)
	}
	ann["kube-bind.io/bound"] = "true"

	obj.SetAnnotations(ann)
}
