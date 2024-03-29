import {
  GoDefaultModelNameConstraints,
  GoFileGenerator,
  ConstrainedDictionaryModel,
} from '@asyncapi/modelina'
import * as path from 'path'
import * as fs from 'fs'

const outputDirPath = path.resolve(__dirname, '../../models')
const schemaFiles = path.resolve(__dirname, '../definitions')

interface FileReadType {
  fileContent: string
  fileName: string
  filePath: string
  asyncapiVersion: string
}

const filteredFiles: FileReadType[] = fs
  .readdirSync(schemaFiles)
  .map((file) => {
    const absPath = path.resolve(schemaFiles, file)
    return {
      fileContent: fs.readFileSync(absPath).toString(),
      fileName: path.basename(file),
      filePath: absPath,
      asyncapiVersion: path.basename(file).replace('-without-$id.json', ''),
    }
  })

async function generateModelsForV2_0() {
  const outputDirForV2_0 = path.resolve(
    __dirname,
    outputDirPath,
    'asyncapi_2_0'
  )
  //const input = versions['2.0.0'] as any;
  const inputFileContent = await fs.promises.readFile(
    path.resolve(__dirname, './asyncapi-2_0.json')
  )
  const input = JSON.parse(String(inputFileContent))
  if (fs.existsSync(outputDirForV2_0)) {
    fs.rmSync(outputDirForV2_0, { recursive: true })
  }
  const generator = new GoFileGenerator({})

  await generator.generateToFiles(input, outputDirForV2_0, {
    packageName: 'models',
  })
}

async function defaultGenerateModels(input: FileReadType, outputDir: string) {
  const inputObj = JSON.parse(String(input.fileContent))
  const outputDirForVersion = path.resolve(
    __dirname,
    outputDir,
    input.asyncapiVersion
  )
  if (fs.existsSync(outputDirForVersion)) {
    fs.rmSync(outputDirForVersion, { recursive: true })
  }
  const generator = new GoFileGenerator({
    presets: [
      {
        struct: {
          field: ({ content, field }) => {
            if (
              field.property instanceof ConstrainedDictionaryModel &&
              field.property.serializationType === 'unwrap'
            ) {
              return `${content} \`json:"-"\``
            }
            return `${content} \`json:"${field.unconstrainedPropertyName}"\``
          },
        },
        enum: {
          self({content,model, renderer}) {
            renderer.dependencyManager.addDependency("encoding/json")
                       const extraCode = `
          
func (op *${model.name}) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesTo${model.name}[v]
	return nil
}

func (op ${model.name}) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          `

            return `${content}\n ${extraCode}` 
          },
        }
      },
    ],
  })

  await generator.generateToFiles(inputObj, outputDirForVersion, {
    packageName: 'models',
  })
}

async function generate() {
  // for (const file of filteredFiles) {
  //   await defaultGenerateModels(file, outputDirPath)
  // }
  const f = filteredFiles[filteredFiles.length - 1]
  await defaultGenerateModels(f, outputDirPath)
}
generate().catch(e => console.log(e))
