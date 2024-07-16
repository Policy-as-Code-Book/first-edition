package model

import (
	"bytes"
	"encoding/json"
	opabundle "github.com/open-policy-agent/opa/bundle"
	"jimmyray.io/opa-bundle-api/pkg/utils"
	"os"
	"strings"
)

const (
	defaultBundleExt          string = ".tar.gz"
	defaultBundleSigningAlg   string = "RS256"
	defaultBundleRegistryPath string = "bundle-registry.json"
)

type BundleRef struct {
	Name     string `json:"bundle-name"`
	Etag     string `json:"bundle-etag"`
	FileName string `json:"bundle-file-name"`
	FilePath string `json:"bundle-file-path"`
	Revision string `json:"bundle-revision"`
}

type BundleReg struct {
	Bundles map[string]BundleRef `json:"bundles"`
}

func (br BundleReg) Json() ([]byte, error) {
	out, err := json.Marshal(br)
	return out, err
}

var RegBundles = BundleReg{}

func BuildBundle() error {
	if SC.Bundles.LoadRegistry {
		if utils.FileExists(defaultBundleRegistryPath) {
			bytes, err := utils.ReadFile(defaultBundleRegistryPath)
			if err != nil {
				utils.Logger.Error().Err(err).Msgf("could not read registry file ", defaultBundleRegistryPath)
				RegBundles.Bundles = make(map[string]BundleRef)
			} else {
				err = json.Unmarshal(bytes, &RegBundles)
				if err != nil {
					utils.Logger.Error().Err(err).Msg("could not unmarshal JSON")
					RegBundles.Bundles = make(map[string]BundleRef)
				}
			}
		}
	} else {
		RegBundles.Bundles = make(map[string]BundleRef)
	}

	j, e := RegBundles.Json()
	if e == nil {
		utils.Logger.Debug().Msgf("Current Bundle Registry: %s", string(j))
	} else {
		utils.Logger.Error().Err(e).Msgf("could not parse bundle registry")
	}

	for _, b := range SC.Bundles.Bundles {
		if b.Build {
			var etag string
			var modules []opabundle.ModuleFile
			var data = make(map[string]interface{})
			bundle := opabundle.Bundle{
				Manifest: opabundle.Manifest{Revision: b.BundleRevision, Roots: &(b.BundleRoots)},
			}

			// Build this bundle
			m, err := utils.FileInfos(b.BundleInDir)
			if err != nil {
				return err
			}

			// Add Manifest to be hashed for etag
			etag += bundle.Manifest.String()

			for k, v := range m {
				if !v.IsDir() && !utils.ArrayContains(b.BundleIgnoreFiles, v.Name()) {

					// Get file bytee
					fileBytes, fileReadError := os.ReadFile(k)
					if fileReadError != nil {
						return fileReadError
					}
					utils.Logger.Debug().Msgf("Processing: %s", k)

					// Only add Rego to modules
					if strings.Contains(v.Name(), ".rego") {
						mf := opabundle.ModuleFile{}
						mf.Raw = fileBytes
						mf.Path = k
						mf.URL = k //Needed for signing

						// Append module file
						modules = append(modules, mf)
					} else if v.Name() == "data.json" {
						err = json.Unmarshal([]byte(fileBytes), &data)
						if err != nil {
							utils.Logger.Error().Err(err).Msg("could not unmarshal JSON")
							return err
						}
					}

					// Collect file content for etag
					etag += string(fileBytes)
				}
			}

			bundle.Data = data
			bundle.Modules = modules
			hash := utils.MD5Hash(etag)
			bundle.Etag = hash

			// Bundle Signing
			if b.BundleSigning.Enable {
				fileBytes, fileReadError := os.ReadFile(b.BundleSigning.SigningKey)
				if fileReadError != nil {
					return fileReadError
				}

				secret := string(fileBytes)
				var alg string
				if b.BundleSigning.SigningAlg == "" {
					alg = defaultBundleSigningAlg
				} else {
					alg = b.BundleSigning.SigningAlg
				}
				if err = bundle.GenerateSignature(opabundle.NewSigningConfig(secret, alg, ""),
					b.BundleSigning.SigningKeyID, false); err != nil {
					return err
				}
			}

			// Write bundle
			var buf bytes.Buffer
			var bundleFilePath string
			var bundleFileName string
			if b.BundleTs {
				// Add bundle timestamp
				bundleFilePath = b.BundleOutDir + "/" + b.BundleFileName + "-" + utils.TsString() + defaultBundleExt
			} else {
				bundleFilePath = b.BundleOutDir + "/" + b.BundleFileName + defaultBundleExt
			}

			if b.BundleTs {
				// Add bundle timestamp
				bundleFileName = b.BundleFileName + "-" + utils.TsString() + defaultBundleExt
			} else {
				bundleFileName = b.BundleFileName + defaultBundleExt
			}

			bundleFilePath = b.BundleOutDir + "/" + bundleFileName

			if err = opabundle.NewWriter(&buf).UseModulePath(true).Write(bundle); err != nil {
				utils.Logger.Error().Err(err).Msgf("could not write bundle %s", bundleFilePath)
				return err
			}

			if utils.FileExists(bundleFilePath) {
				err = os.Remove(bundleFilePath)
				if err != nil {
					utils.Logger.Error().Err(err).Msgf("could not delete file ", bundleFilePath)
				}
			}

			if err = os.WriteFile(bundleFilePath, buf.Bytes(), 0420); err != nil {
				utils.Logger.Error().Err(err).Msgf("could not write bundle file ", bundleFilePath)
				return err
			}

			// Register bundle
			br := BundleRef{
				Name:     b.BundleName,
				Etag:     hash,
				FilePath: bundleFilePath,
				FileName: bundleFileName,
				Revision: b.BundleRevision,
			}
			RegBundles.Bundles[br.FileName] = br
		}
	}

	if SC.Bundles.PersistRegistry {
		if utils.FileExists(defaultBundleRegistryPath) {
			err := os.Remove(defaultBundleRegistryPath)
			if err != nil {
				utils.Logger.Error().Err(err).Msgf("could not delete file %s", defaultBundleRegistryPath)
			}
		}

		// Write registry JSON
		bytes, err := json.Marshal(RegBundles)
		if err != nil {
			utils.Logger.Error().Err(err).Msg("could not marshal JSON from bundle registry")
			return err
		}

		err = utils.WriteFile(defaultBundleRegistryPath, bytes)
		if err != nil {
			utils.Logger.Error().Err(err).Msg("could not write JSON bundle registry")
			return err
		}
	}

	return nil
}

func PurgeBundles() error {
	// Purge bundle files
	files, err := os.ReadDir(SC.Bundles.BundleOutDir)
	if err != nil {
		utils.Logger.Error().Err(err).Msgf("could not list bundle dir %s", SC.Bundles.BundleOutDir)
	} else {
		for _, f := range files {
			if strings.Contains(f.Name(), ".tar.gz") {
				err = os.Remove(SC.Bundles.BundleOutDir + "/" + f.Name())
				if err != nil {
					utils.Logger.Error().Err(err).Msgf("could not delete file %s/%s", SC.Bundles.BundleOutDir, f.Name())
				}
			}
		}
	}

	//Purge bundle registry
	if utils.FileExists(defaultBundleRegistryPath) {
		err = os.Remove(defaultBundleRegistryPath)
		if err != nil {
			utils.Logger.Error().Err(err).Msgf("could not delete %s", defaultBundleRegistryPath)
			return err
		}
	}

	return nil
}
