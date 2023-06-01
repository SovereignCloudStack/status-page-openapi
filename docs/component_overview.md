# Component Overview

This represents a part of the decission process related to the overall structure the API wants to represent.

```mermaid
C4Component
  title status-page-openapi

  Container_Boundary(impacts, "Impacts") {
    Component(impactType, ImpactType, "ID, displayName, description")
    Component(impact, Impact, "type, reference")
    Component(impactList, ImpactList, "[]Impact")

    Rel(impact, impactType, "has")
    Rel(impactList, impact, "lists")

    UpdateElementStyle(impact, $bgColor="green")
    UpdateElementStyle(impactList, $bgColor="green")

    UpdateRelStyle(impact, impactType, "green", "green", $offsetY="-10")
    UpdateRelStyle(impactList, impact, "green", "green", $offsetY="-10")
  }

  Container_Boundary(incidents, "Incidents") {
    Component(incidentUpdate, IncidentUpdate, "order, displayName, description, createdAt")
    Component(incident, Incident, "ID, displayName, description, updates, affects, beganAt, endedAt, Phase")

    Rel(incident, incidentUpdate, "contains")
  }

  Container_Boundary(phases, "Phases") {
    Component(phaseReference, PhaseReference, "Phase, order, generation")
    Component(phase, Phase, "", "it is just a name")
    Component(phaseList, PhaseList, "generation, []Phase")

    Rel(phaseList, phase, "lists")
    Rel(phaseReference, phase, "references")
    Rel(phaseReference, phaseList, "references")

    UpdateElementStyle(phaseReference, $bgColor="green")
    UpdateElementStyle(phaseList, $bgColor="green")

    UpdateRelStyle(phaseList, phase, "green", "green", $offsetY="5")
    UpdateRelStyle(phaseReference, phase, "green", "green", $offsetX="-30", $offsetY="10")
    UpdateRelStyle(phaseReference, phaseList, "green", "green", $offsetY="5")

  }

  Container_Boundary(components, "Components") {
    Component(component, Component, "ID, displayName, Labels, affectedBy")
    Component(labels, Labels, "", "Key value pairs")

    Rel(component, labels, "contains")
  }

  %% global relations %%
  %% deprecated %%
  Rel(incident, impactType, "has")
  BiRel(component, incident, "affects / affected by")

  UpdateRelStyle(incident, impactType, "red", "red")
  UpdateRelStyle(component, incident, "red", "red", $offsetX="10")

  %% new %%
  Rel(incident, impactList, "affects")
  Rel(component, impactList, "affected by")
  Rel(incident, phaseReference, "has")
  Rel(impact, component, "references", "from Incident")
  Rel(impact, incident, "references", "from Component")

  UpdateRelStyle(incident, impactList, "green", "green")
  UpdateRelStyle(component, impactList, "green", "green")
  UpdateRelStyle(incident, phaseReference, "green", "green")
  UpdateRelStyle(impact, component, "green", "green")
  UpdateRelStyle(impact, incident, "green", "green")

```

Color meaning:

- Blue: Existing structure
- Red: Deprecated / removed
- Green: New
