export * from './component.service';
import { ComponentService } from './component.service';
export * from './component.serviceInterface';
export * from './impact.service';
import { ImpactService } from './impact.service';
export * from './impact.serviceInterface';
export * from './incident.service';
import { IncidentService } from './incident.service';
export * from './incident.serviceInterface';
export * from './phase.service';
import { PhaseService } from './phase.service';
export * from './phase.serviceInterface';
export const APIS = [ComponentService, ImpactService, IncidentService, PhaseService];
