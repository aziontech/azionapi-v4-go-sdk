import sys
import yaml
import os
import argparse

# Define the structure mappings for renaming
STRUCTURE_MAPPINGS = {
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineNoArgsRequest": "EdgeApplicationBehaviorNoArgs",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineStringRequest": "EdgeApplicationBehaviorString",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineRunFunctionRequest": "EdgeApplicationBehaviorInteger",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest": "EdgeApplicationBehaviorObject"
}

# Define the mapping changes for the discriminator mapping
MAPPING_CHANGES = {
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineNoArgsRequest": "EdgeApplicationBehaviorNoArgs",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineStringRequest": "EdgeApplicationBehaviorString",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineFilterRequestCookieRequest": "EdgeApplicationBehaviorString",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineAddHeaderRequest": "EdgeApplicationBehaviorString",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineRewriteRequestRequest": "EdgeApplicationBehaviorString",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineFilterHeaderRequest": "EdgeApplicationBehaviorString",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineAddRequestCookieRequest": "EdgeApplicationBehaviorString",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest": "EdgeApplicationBehaviorObject",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineRunFunctionRequest": "EdgeApplicationBehaviorInteger",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineSetOriginRequest": "EdgeApplicationBehaviorInteger",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineSetConnectorRequest": "EdgeApplicationBehaviorInteger",
    "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineSetCachePolicyRequest": "EdgeApplicationBehaviorInteger"
}

def rename_structures_in_yaml(file_path):
    """Rename structures in the YAML file and update references."""
    with open(file_path, "r") as file:
        content = file.read()
    
    # Load YAML while preserving order
    data = yaml.safe_load(content)
    
    # Check if the file has components and schemas
    if "components" not in data or "schemas" not in data["components"]:
        print(f"Skipping {file_path}: No components/schemas found")
        return
    
    # Create new schema definitions with the new names
    for old_name, new_name in STRUCTURE_MAPPINGS.items():
        if old_name in data["components"]["schemas"]:
            # Copy the schema with the new name
            data["components"]["schemas"][new_name] = data["components"]["schemas"][old_name].copy()
            
            # Update the title if it exists
            if "title" in data["components"]["schemas"][new_name]:
                data["components"]["schemas"][new_name]["title"] = new_name
    
    # Update the EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest schema to use OneOf
    if "EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest" in data["components"]["schemas"]:
        # Create the OneOf structure
        data["components"]["schemas"]["EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest"]["oneOf"] = [
            {"$ref": f"#/components/schemas/{new_name}"} for new_name in STRUCTURE_MAPPINGS.values()
        ]
        
        # Remove any existing allOf or anyOf if present
        if "allOf" in data["components"]["schemas"]["EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest"]:
            del data["components"]["schemas"]["EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest"]["allOf"]
        if "anyOf" in data["components"]["schemas"]["EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest"]:
            del data["components"]["schemas"]["EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest"]["anyOf"]
    
    # Update all references to the old structure names
    update_references_in_yaml(data, STRUCTURE_MAPPINGS)
    
    # Write the updated YAML back to the file
    with open(file_path, "w") as file:
        yaml.dump(data, file, default_flow_style=False, sort_keys=False)
    
    print(f"✅ Updated structure names in {file_path}")
    
    # Now update the mapping object
    update_mapping_in_yaml(file_path)

def update_references_in_yaml(data, mappings):
    """Recursively update references in the YAML data."""
    if isinstance(data, dict):
        for key, value in list(data.items()):
            if key == "$ref" and isinstance(value, str):
                for old_name, new_name in mappings.items():
                    if f"#/components/schemas/{old_name}" in value:
                        data[key] = value.replace(f"#/components/schemas/{old_name}", f"#/components/schemas/{new_name}")
            else:
                update_references_in_yaml(value, mappings)
    elif isinstance(data, list):
        for item in data:
            update_references_in_yaml(item, mappings)

def update_mapping_in_yaml(file_path):
    """Update the mapping object in EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest schema."""
    with open(file_path, "r") as file:
        content = file.read()
    
    # Load YAML while preserving order
    data = yaml.safe_load(content)
    
    # Check if the file has components and schemas
    if "components" not in data or "schemas" not in data["components"]:
        return
    
    # Check if EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest exists
    if "EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest" not in data["components"]["schemas"]:
        return
    
    # Get the schema
    schema = data["components"]["schemas"]["EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest"]
    
    # Check if it has a discriminator with mapping
    if "discriminator" not in schema or "mapping" not in schema["discriminator"]:
        return
    
    # Update the mapping
    mapping = schema["discriminator"]["mapping"]
    updated = False
    
    for key, value in mapping.items():
        for old_name, new_name in MAPPING_CHANGES.items():
            if old_name in value:
                mapping[key] = value.replace(old_name, new_name)
                updated = True
    
    if updated:
        # Write the updated YAML back to the file
        with open(file_path, "w") as file:
            yaml.dump(data, file, default_flow_style=False, sort_keys=False)
        print(f"✅ Updated mapping in {file_path}")

def process_yaml_files(directory):
    """Process all YAML files in the given directory."""
    yaml_files = []
    
    # Find all YAML files
    for root, _, files in os.walk(directory):
        for file in files:
            if file.endswith((".yaml", ".yml")):
                yaml_files.append(os.path.join(root, file))
    
    # Process each YAML file
    for yaml_file in yaml_files:
        try:
            rename_structures_in_yaml(yaml_file)
        except Exception as e:
            print(f"Error processing {yaml_file}: {e}")

def main():
    parser = argparse.ArgumentParser(description="Rename structures and update mappings in YAML files")
    parser.add_argument("directory", help="Directory containing YAML files to process")
    args = parser.parse_args()
    
    process_yaml_files(args.directory)
    print("✅ Structure renaming and mapping update completed successfully.")

if __name__ == "__main__":
    main()
