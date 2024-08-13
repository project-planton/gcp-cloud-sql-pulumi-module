package pkg

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpcloudsql"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	Input     *gcpcloudsql.GcpCloudSqlStackInput
	GcpLabels map[string]string
}

func (s *ResourceStack) Resources(ctx *pulumi.Context) error {
	//create gcp provider using the credentials from the input
	_, err := pulumigoogleprovider.Get(ctx, s.Input.GcpCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}
	return nil
}
