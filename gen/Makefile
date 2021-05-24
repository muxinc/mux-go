# per-SDK stuff
GENERATOR_TYPE=go

###### helper variables ######
# this is in case we want to go to Docker or something other than running yarn
# for OAS in the future. also less typing.
OAS_CLI=yarn run -s oas
CONFIG_PATH=./generator-config.json
OUTPUT_DIR=..
TEMPLATE_DIR=./templates
GENERATED_TEMPLATE_DIR=./tmp/oas-base-templates

ensure:
	yarn install

clean:
	rm -rf tmp

clean-products:
	rm -rf "${OUTPUT_DIR}/lib"
	rm -rf "${OUTPUT_DIR}/docs"

build: ensure clean-products
	${OAS_CLI} generate -g "${GENERATOR_TYPE}" -c "${CONFIG_PATH}" -t "${TEMPLATE_DIR}" -o "${OUTPUT_DIR}" -i "${OAS_SPEC_PATH}"

config-help: ensure
	${OAS_CLI} config-help -g "${GENERATOR_TYPE}"

template-diff: ensure
	${OAS_CLI} author template -g "${GENERATOR_TYPE}" -o "${GENERATED_TEMPLATE_DIR}"
	diff "${TEMPLATE_DIR}" "${GENERATED_TEMPLATE_DIR}" > ./tmp/template-changes.diff
