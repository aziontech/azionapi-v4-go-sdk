#!/usr/bin/env python3
import sys
import yaml
import re

def replace_strings_in_yaml(content):
    """
    Replace specific strings in the YAML content.
    """
    replacements = {
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineNoArgsRequest": "EdgeApplicationRequestPhaseBehaviorWithoutArgsRequest",
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineStringRequest": "EdgeApplicationRequestPhaseBehaviorWithArgsRequest",
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest": "EdgeApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest",
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineNoArgs": "EdgeApplicationRequestPhaseBehaviorWithoutArgs",
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineString": "EdgeApplicationRequestPhaseBehaviorWithArgs",
        "EdgeApplicationRuleEngineRequestPhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroups": "EdgeApplicationRequestPhaseBehaviorCaptureMatchGroups",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineResponseNoArgsRequest": "EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineResponseStringRequest": "EdgeApplicationResponsePhaseBehaviorWithArgsRequest",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroupsRequest":  "EdgeApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineResponseNoArgs": "EdgeApplicationResponsePhaseBehaviorWithoutArgs",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineResponseString": "EdgeApplicationResponsePhaseBehaviorWithArgs",
        "EdgeApplicationRuleEngineResponsePhaseBehaviorsEdgeApplicationRuleEngineCaptureMatchGroups":  "EdgeApplicationResponsePhaseBehaviorCaptureMatchGroups"
        
    }
    
    # Perform string replacements
    for old_str, new_str in replacements.items():
        content = content.replace(old_str, new_str)
    
    print("✅ Successfully replaced strings in YAML")
    return content


def modify_request_phase_behaviors(data):
    """
    Modify request phase behaviors
    """
    
    allowed_refs = [
        {'$ref': '#/components/schemas/EdgeApplicationRequestPhaseBehaviorWithoutArgs'},
        {'$ref': '#/components/schemas/EdgeApplicationRequestPhaseBehaviorWithArgs'},
        # {'$ref': '#/components/schemas/EdgeApplicationRequestPhaseBehaviorCaptureMatchGroups'}
    ]
    
    allowed_refs_request = [
        {'$ref': '#/components/schemas/EdgeApplicationRequestPhaseBehaviorWithoutArgsRequest'},
        {'$ref': '#/components/schemas/EdgeApplicationRequestPhaseBehaviorWithArgsRequest'},
        {'$ref': '#/components/schemas/EdgeApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest'}
    ]

    # Update oneOfs
    data['components']['schemas']['EdgeApplicationRuleEngineRequestPhaseBehaviors']['oneOf'] = allowed_refs
    data['components']['schemas']['EdgeApplicationRuleEngineRequestPhaseBehaviorsRequest']['oneOf'] = allowed_refs_request
    
    # Remove additionalProperties from withoutArgs struct to prevent problems
    if data['components']['schemas']['EdgeApplicationRequestPhaseBehaviorWithoutArgs']['additionalProperties']:
        data['components']['schemas']['EdgeApplicationRequestPhaseBehaviorWithoutArgs']['additionalProperties'] = False
    
    # Adapt allOf behavior schema
    data['components']['schemas']['EdgeApplicationRequestPhaseBehaviorWithoutArgs']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineNoArgs'}] 
    data['components']['schemas']['EdgeApplicationRequestPhaseBehaviorWithArgs']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineString'}]
    data['components']['schemas']['EdgeApplicationRequestPhaseBehaviorCaptureMatchGroups']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineCaptureMatchGroupsRequest'}]
    data['components']['schemas']['EdgeApplicationRequestPhaseBehaviorWithoutArgsRequest']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineNoArgs'}] 
    data['components']['schemas']['EdgeApplicationRequestPhaseBehaviorWithArgsRequest']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineString'}]
    data['components']['schemas']['EdgeApplicationRequestPhaseBehaviorCaptureMatchGroupsRequest']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineCaptureMatchGroupsRequest'}]
    
    # Adapt value field validation
    data['components']['schemas']['EdgeApplicationRuleEngineStringAttributes']['properties']['value'] = {
        "oneOf": [
            { "type": "string" },
            { "type": "integer", "format": "int64" }
        ]
    }
    
    # Remove Required Fields
    if 'required' in data['components']['schemas']['EdgeApplicationRuleEngineStringAttributes']:
        del data['components']['schemas']['EdgeApplicationRuleEngineStringAttributes']['required']

    print("✅ Successfully modified Request Phase Behaviors")


def modify_response_phase_behaviors(data):
    """
    Modify response phase behaviors
    """
    
    allowed_refs = [
        {'$ref': '#/components/schemas/EdgeApplicationResponsePhaseBehaviorWithoutArgs'},
        {'$ref': '#/components/schemas/EdgeApplicationResponsePhaseBehaviorWithArgs'},
        # {'$ref': '#/components/schemas/EdgeApplicationResponsePhaseBehaviorCaptureMatchGroups'}
    ]
    
    allowed_refs_request = [
        {'$ref': '#/components/schemas/EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest'},
        {'$ref': '#/components/schemas/EdgeApplicationResponsePhaseBehaviorWithArgsRequest'},
        {'$ref': '#/components/schemas/EdgeApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest'}
    ]

    # Update oneOfs
    data['components']['schemas']['EdgeApplicationRuleEngineResponsePhaseBehaviors']['oneOf'] = allowed_refs
    data['components']['schemas']['EdgeApplicationRuleEngineResponsePhaseBehaviorsRequest']['oneOf'] = allowed_refs_request
    
    # Remove additionalProperties from withoutArgs struct to prevent problems
    if data['components']['schemas']['EdgeApplicationResponsePhaseBehaviorWithoutArgs']['additionalProperties']:
        data['components']['schemas']['EdgeApplicationResponsePhaseBehaviorWithoutArgs']['additionalProperties'] = False
    
    # Adapt allOf behavior schema
    data['components']['schemas']['EdgeApplicationResponsePhaseBehaviorWithoutArgs']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineNoArgs'}] 
    data['components']['schemas']['EdgeApplicationResponsePhaseBehaviorWithArgs']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineString'}]
    data['components']['schemas']['EdgeApplicationResponsePhaseBehaviorCaptureMatchGroups']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineCaptureMatchGroups'}]
    data['components']['schemas']['EdgeApplicationResponsePhaseBehaviorWithoutArgsRequest']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineNoArgs'}] 
    data['components']['schemas']['EdgeApplicationResponsePhaseBehaviorWithArgsRequest']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineString'}]
    data['components']['schemas']['EdgeApplicationResponsePhaseBehaviorCaptureMatchGroupsRequest']['allOf'] = [{'$ref': '#/components/schemas/EdgeApplicationRuleEngineCaptureMatchGroups'}]


    print("✅ Successfully modified Response Phase Behaviors")

    
def process_yaml_file(file_path):
    """
    Process a YAML file by loading it, replacing strings, and modifying component schema.
    """
    try:
        # Read the file
        with open(file_path, 'r') as file:
            yaml_content = file.read()
        
        # Replace strings
        modified_content = replace_strings_in_yaml(yaml_content)
        
        # Load YAML to modify component schema
        data = yaml.safe_load(modified_content)
        
        # Modify behaviors
        modify_request_phase_behaviors(data)
        modify_response_phase_behaviors(data)
        
        # Write back to file
        with open(file_path, 'w') as file:
            yaml.dump(data, file, default_flow_style=False, sort_keys=False)
        
        print(f"✅ Successfully processed {file_path}")
        
    except Exception as e:
        print(f"❌ Error processing {file_path}: {str(e)}")
        sys.exit(1)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python replace_strings.py <path-to-yaml-file>")
        sys.exit(1)
    
    yaml_file_path = sys.argv[1]
    process_yaml_file(yaml_file_path)
