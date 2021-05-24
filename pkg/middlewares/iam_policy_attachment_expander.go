package middlewares

import (
	"fmt"

	"github.com/cloudskiff/driftctl/pkg/resource"
	resourceaws "github.com/cloudskiff/driftctl/pkg/resource/aws"
)

// Split Policy attachment when there is multiple user and groups and generate a repeatable id
type IamPolicyAttachmentExpander struct {
	resourceFactory resource.ResourceFactory
}

func NewIamPolicyAttachmentExpander(resourceFactory resource.ResourceFactory) IamPolicyAttachmentExpander {
	return IamPolicyAttachmentExpander{
		resourceFactory,
	}
}

func (m IamPolicyAttachmentExpander) Execute(remoteResources, resourcesFromState *[]resource.Resource) error {
	var newStateResources = make([]resource.Resource, 0)

	for _, stateResource := range *resourcesFromState {
		// Ignore all resources other than policy attachment
		if stateResource.TerraformType() != resourceaws.AwsIamPolicyAttachmentResourceType {
			newStateResources = append(newStateResources, stateResource)
			continue
		}

		policyAttachment := stateResource.(*resource.AbstractResource)

		newStateResources = append(newStateResources, m.expand(policyAttachment)...)
	}

	var newRemoteResources = make([]resource.Resource, 0)

	for _, remoteResource := range *remoteResources {
		// Ignore all resources other than policy attachment
		if remoteResource.TerraformType() != resourceaws.AwsIamPolicyAttachmentResourceType {
			newRemoteResources = append(newRemoteResources, remoteResource)
			continue
		}

		policyAttachment := remoteResource.(*resource.AbstractResource)

		newRemoteResources = append(newRemoteResources, m.expand(policyAttachment)...)
	}

	*resourcesFromState = newStateResources
	*remoteResources = newRemoteResources

	return nil
}

func (m IamPolicyAttachmentExpander) expand(policyAttachment *resource.AbstractResource) []resource.Resource {

	arn, _ := policyAttachment.Attrs.Get("policy_arn")
	user, exist := policyAttachment.Attrs.Get("user")
	if exist {
		user := user.(string)
		newAttachment := m.resourceFactory.CreateAbstractResource(
			resourceaws.AwsIamPolicyAttachmentResourceType,
			fmt.Sprintf("%s-%s", user, arn),
			map[string]interface{}{
				"users": []string{user},
			},
		)
		return []resource.Resource{newAttachment}
	}

	role, exist := policyAttachment.Attrs.Get("role")
	if exist {
		role := role.(string)
		newAttachment := m.resourceFactory.CreateAbstractResource(
			resourceaws.AwsIamPolicyAttachmentResourceType,
			fmt.Sprintf("%s-%s", role, arn),
			map[string]interface{}{
				"roles": []string{role},
			},
		)
		return []resource.Resource{newAttachment}
	}

	var newResources []resource.Resource

	users := (*policyAttachment.Attrs)["users"]
	if users != nil {
		// we create one attachment per user
		for _, user := range users.([]interface{}) {
			user := user.(string)
			newAttachment := m.resourceFactory.CreateAbstractResource(
				resourceaws.AwsIamPolicyAttachmentResourceType,
				fmt.Sprintf("%s-%s", user, (*policyAttachment.Attrs)["policy_arn"]),
				map[string]interface{}{
					"users": []string{user},
				},
			)
			newResources = append(newResources, newAttachment)
		}
	}

	roles := (*policyAttachment.Attrs)["roles"]
	if roles != nil {
		// we create one attachment per role
		for _, role := range roles.([]interface{}) {
			role := role.(string)
			newAttachment := m.resourceFactory.CreateAbstractResource(
				resourceaws.AwsIamPolicyAttachmentResourceType,
				fmt.Sprintf("%s-%s", role, (*policyAttachment.Attrs)["policy_arn"]),
				map[string]interface{}{
					"roles": []string{role},
				},
			)
			newResources = append(newResources, newAttachment)
		}
	}
	return newResources
}