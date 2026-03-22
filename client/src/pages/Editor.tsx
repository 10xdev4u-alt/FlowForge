import React, { useCallback } from 'react';
import ReactFlow, { 
  Background, 
  Controls, 
} from 'reactflow';
import 'reactflow/dist/style.css';
import { useWorkflowStore } from '../store/workflowStore';
import { TriggerNode } from '../components/nodes/TriggerNode';
import { ActionNode } from '../components/nodes/ActionNode';

const nodeTypes = {
  trigger: TriggerNode,
  webhook: TriggerNode,
  action: ActionNode,
  http: ActionNode,
};

export const Editor = () => {
  const { nodes, edges, onNodesChange, onEdgesChange, onConnect } = useWorkflowStore();

  return (
    <div className="w-full h-screen">
      <ReactFlow
        nodes={nodes}
        edges={edges}
        nodeTypes={nodeTypes}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        onConnect={onConnect}
      >
        <Controls />
        <Background />
      </ReactFlow>
    </div>
  );
};
