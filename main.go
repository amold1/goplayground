package main

import (
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

func retryOnErrorExample() {
	count := 0
	err := retry.OnError(
		retry.DefaultRetry,
		func(err error) bool {
			if err.Error() == "MyError" {
				return true
			}
			return false
		},
		func() error {
			count++
			fmt.Println(count)
			if count > 4 {
				fmt.Println(count)
				return fmt.Errorf("What")
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
}

// ClusterClose persists the linodecluster configuration and status.
func newRetryOnErrorExample() error {
	var err error
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		return apierrors.NewApplyConflict([]v1.StatusCause{}, "myconflict")
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// v8GetRecord()
	// if err := setRecord("test.akafn.com", "A", "1.1.1.4"); err != nil {
	// 	log.Fatal(err)
	// }
	// result, err := getRecord("dodo.akafn.com", "A")
	// if strings.Contains(err.Error(), "sundar") {
	// 	fmt.Println("Jai Gopal")
	// } else {
	// 	fmt.Println(result)
	// }
	// GetKey()
	// getAllRecordSets()
	// interactWithLinodeGo()
	// var data = Cluster{
	// 	Name:   "abir",
	// 	IPList: []string{"ip1", "ip2", "ip3"},
	// }

	// fileIO(data)

	fmt.Println(newRetryOnErrorExample())
}
