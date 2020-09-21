package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraform(t *testing.T){
	t.Parallel()

	const projectId string  = "ryan-cicd"
	instanceName := fmt.Sprintf("gcp-terratest-test-%s", strings.ToLower(random.UniqueId()))

	terraformOptions := &terraform.Options{
		TerraformDir: "../example2/",

		Vars: map[string]interface{}{
			"instance_name": instanceName,
		},

		EnvVars: map[string]string{
			"GOOGLE_CLOUD_PROJECT": projectId,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
}