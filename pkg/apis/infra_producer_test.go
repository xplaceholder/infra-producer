package apis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kun-lun/infra-producer/pkg/apis"

	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kun-lun/infra-producer/application"
	"github.com/kun-lun/infra-producer/storage"
)

var _ = Describe("InfraProducer", func() {
	var (
		artifact      string
		stateDir      string
		globalConfig  application.GlobalConfiguration
		azureConfig   storage.Azure
		state         storage.State
		infraProducer InfraProducer
	)

	BeforeEach(func() {
		artifact = "artifacts/lamp.yml"
		stateDir = filepath.Join("/tmp/", "kunlun")
		os.MkdirAll(stateDir, os.ModePerm)
		globalConfig = application.GlobalConfiguration{
			StateDir: stateDir,
			Debug:    true,
			Name:     "kunlun",
		}
		azureConfig = storage.Azure{
			Region:         "eastus2",
			Environment:    "AzureCloud",
			SubscriptionID: "5ef1a6b6-cbfc-4e9f-bdae-1afb35682e7e",
			TenantID:       "262f9b5b-376d-4c2e-ba1e-35372fb45d7f",
			ClientID:       "d6d209e8-ec28-4b62-9ff6-826a3e09f6e1",
			ClientSecret:   "e97b306c-2941-459c-97f3-2fba79c71b76",
		}
		state = storage.State{
			Version:        1,
			ID:             "bd73f71c-345a-434d-98bc-4454e6804a00",
			EnvName:        "kunlun",
			Azure:          azureConfig,
			LatestTFOutput: "",
		}
		infraProducer = NewInfraProducer(globalConfig)
	})

	AfterEach(func() {
		os.RemoveAll(stateDir)
	})

	Describe("foo", func() {
		Context("bar", func() {
			It("should setup and apply terraform", func() {
				b, _ := ioutil.ReadFile(artifact)
				Expect(infraProducer.Setup(b, state)).To(BeNil())
				Expect(infraProducer.Apply(state)).To(BeNil())
			})
		})
	})
})
