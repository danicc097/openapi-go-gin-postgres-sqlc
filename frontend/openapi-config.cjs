/** @type {import("@rtk-query/codegen-openapi").ConfigFile} */
const config = {
  schemaFile: '../openapi.yaml',
  apiFile: './src/redux/slices/emptyApi.ts',
  apiImport: 'emptyInternalApi',
  outputFile: './src/redux/slices/gen/internalApi.ts',
  exportName: 'internalApi',
  hooks: true,
  tag: true,
  argSuffix: 'Args',
  responseSuffix: 'Res',
  flattenArg: true,
}

module.exports = config
