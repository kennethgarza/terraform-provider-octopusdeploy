---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "octopusdeploy_space Data Source - terraform-provider-octopusdeploy"
subcategory: ""
description: |-
  Provides information about an existing space.
---

# octopusdeploy_space (Data Source)

Provides information about an existing space.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The name of this resource.

### Read-Only

- **description** (String) The description of this space.
- **id** (String) The unique ID for this resource.
- **is_default** (Boolean) Specifies if this space is the default space in Octopus.
- **is_task_queue_stopped** (Boolean) Specifies the status of the task queue for this space.
- **space_managers_team_members** (List of String) A list of user IDs designated to be managers of this space.
- **space_managers_teams** (List of String) A list of team IDs designated to be managers of this space.

