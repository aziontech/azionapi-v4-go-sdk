import sys
import yaml

def infer_type_from_enum(enum_values):
    """Try to infer the type from the first value in an enum list."""
    if not enum_values:
        return None

    first_value = enum_values[0]

    if isinstance(first_value, int):
        return "integer"
    elif isinstance(first_value, str):
        return "string"
    else:
        return None

def process_schema(schema, parent_key="root"):
    """Recursively process schema definitions in the OpenAPI document."""
    if isinstance(schema, dict):
        inferred_type = schema.get("type")

        # If there is an enum but no type, infer type and then delete enum
        if "enum" in schema:
            if not inferred_type:
                inferred_type = infer_type_from_enum(schema["enum"])
                if inferred_type:
                    schema["type"] = inferred_type
                else:
                    print(f"⚠️ Manual action needed: Unable to infer type for '{parent_key}' (was enum, now missing type)")
            del schema["enum"]

        # Ensure integer types have format int64 if format is missing
        if schema.get("type") == "integer" and "format" not in schema:
            schema["format"] = "int64"

        # Remove "additionalProperties: false" from string or integer objects
        if schema.get("type") in ["string", "integer"] and "additionalProperties" in schema:
            del schema["additionalProperties"]

        # Remove default field if present
        if "default" in schema:
            del schema["default"]

        #change default_args to an empty dictionary
        if "default_args" in schema:
            schema["default_args"] = {}
        
        # Overwrite json_args with an empty dictionary
        if "json_args" in schema:
            schema["json_args"] = {}

        # Recursively process all values
        for key, value in schema.items():
            process_schema(value, parent_key=f"{parent_key}.{key}")
    elif isinstance(schema, list):
        for i, item in enumerate(schema):
            process_schema(item, parent_key=f"{parent_key}[{i}]")

def modify_openapi_yaml(file_path):
    """Load YAML, apply transformations, and save changes."""
    with open(file_path, "r") as file:
        data = yaml.safe_load(file)

    if "components" in data and "schemas" in data["components"]:
        for schema_name, schema in data["components"]["schemas"].items():
            process_schema(schema, parent_key=schema_name)

    if "paths" in data:
        for path, methods in data["paths"].items():
            for method, details in methods.items():
                if isinstance(details, dict):
                    process_schema(details, parent_key=f"{path}.{method}")

    with open(file_path, "w") as file:
        yaml.dump(data, file, default_flow_style=False, sort_keys=False)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python modify_openapi.py <path-to-openapi.yaml>")
        sys.exit(1)

    modify_openapi_yaml(sys.argv[1])
    print("✅ YAML modifications applied successfully.")
