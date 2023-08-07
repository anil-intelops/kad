package credential

import (
	"context"
	"fmt"

	"github.com/intelops/go-common/credentials"
	"github.com/pkg/errors"
)

const (
	clusterCertEntity = "client-cert"
)

func GetServiceUserCredential(ctx context.Context, svcEntity, userName string) (cred credentials.ServiceCredential, err error) {
	credReader, err := credentials.NewCredentialReader(ctx)
	if err != nil {
		err = errors.WithMessage(err, "error in initializing credential reader")
		return
	}

	cred, err = credReader.GetServiceUserCredential(context.Background(), svcEntity, userName)
	if err != nil {
		err = errors.WithMessagef(err, "error in reading credential for %s/%s", svcEntity, userName)
	}
	return
}

func GetClusterCerts(ctx context.Context, orgID, clusterName string) (cred credentials.CertificateData, err error) {
	credReader, err := credentials.NewCredentialReader(ctx)
	if err != nil {
		err = errors.WithMessage(err, "error in initializing credential reader")
		return
	}

	certIndetifier := getClusterCertIndentifier(orgID, clusterName)
	cred, err = credReader.GetCertificateData(context.Background(), clusterCertEntity, certIndetifier)
	if err != nil {
		err = errors.WithMessagef(err, "error in reading cert for %s/%s", clusterCertEntity, certIndetifier)
	}
	return
}

func GetGenericCredential(ctx context.Context, entityName, credIndentifer string) (cred map[string]string, err error) {
	credReader, err := credentials.NewCredentialReader(ctx)
	if err != nil {
		err = errors.WithMessage(err, "error in initializing credential reader")
		return
	}

	cred, err = credReader.GetCredential(context.Background(), credentials.GenericCredentialType, entityName, credIndentifer)
	if err != nil {
		err = errors.WithMessagef(err, "error in reading cred for %s/%s", clusterCertEntity, credIndentifer)
	}
	return
}

func PutServiceUserCredential(ctx context.Context, svcEntity, userIdentifer, userName, password string) error {
	credAdmin, err := credentials.NewCredentialAdmin(ctx)
	if err != nil {
		return errors.WithMessage(err, "error in initializing credential admin")
	}

	err = credAdmin.PutServiceUserCredential(context.Background(), svcEntity, userIdentifer,
		credentials.ServiceCredential{
			UserName: userName,
			Password: password,
		})
	if err != nil {
		return errors.WithMessagef(err, "error in put service cred for %s/%s", svcEntity, userIdentifer)
	}
	return nil
}

func PutClusterCerts(ctx context.Context, orgID, clusterName, clientCAChainData, clientKeyData, clientCertData string) error {
	credAdmin, err := credentials.NewCredentialAdmin(ctx)
	if err != nil {
		return errors.WithMessage(err, "error in initializing credential admin")
	}

	certIndetifier := getClusterCertIndentifier(orgID, clusterName)
	err = credAdmin.PutCertificateData(context.Background(), clusterCertEntity, certIndetifier,
		credentials.CertificateData{
			CACert: clientCAChainData,
			Key:    clientKeyData,
			Cert:   clientCertData,
		})
	if err != nil {
		return errors.WithMessagef(err, "error in put cert for %s/%s", clusterCertEntity, certIndetifier)
	}
	return nil
}

func DeleteClusterCerts(ctx context.Context, orgID, clusterName string) (err error) {
	credReader, err := credentials.NewCredentialAdmin(ctx)
	if err != nil {
		err = errors.WithMessage(err, "error in initializing credential admin")
		return
	}

	certIndetifier := getClusterCertIndentifier(orgID, clusterName)
	err = credReader.DeleteCertificateData(context.Background(), clusterCertEntity, certIndetifier)
	if err != nil {
		err = errors.WithMessagef(err, "error in delete cert for %s/%s", clusterCertEntity, certIndetifier)
	}
	return
}

func getClusterCertIndentifier(orgID, clusterName string) string {
	return fmt.Sprintf("%s:%s", orgID, clusterName)
}
