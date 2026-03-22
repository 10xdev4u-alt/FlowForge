import React, { useCallback } from 'react';
import ReactFlow, { 
  Background, 
  Controls, 
} from 'reactflow';
import 'reactflow/dist/style.css';
import { useWorkflowStore } from '../store/workflowStore';
import { TriggerNode } from '../components/nodes/TriggerNode';
import { ActionNode } from '../components/nodes/ActionNode';

import { useParams, useNavigate } from 'react-router-dom';
import api from '../api/axios';
import { Button } from '../components/ui/Button';

const nodeTypes = {
  trigger: TriggerNode,
  webhook: TriggerNode,
  action: ActionNode,
  http: ActionNode,
};

export const Editor = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { nodes, edges, onNodesChange, onEdgesChange, onConnect, setNodes, setEdges } = useWorkflowStore();

  const onSave = async () => {
    try {
      await api.post(`/workflows/${id}/canvas`, { nodes, edges });
      alert('Workflow saved!');
    } catch (error) {
      alert('Failed to save workflow');
    }
  };

  return (
    <div className="w-full h-screen relative">
      <div className="absolute top-4 left-4 z-10 flex items-center space-x-4">
        <Button variant="secondary" onClick={() => navigate('/')}>Back</Button>
        <h1 className="text-xl font-bold bg-white px-4 py-2 rounded shadow">Workflow Editor</h1>
      </div>
      <div className="absolute top-4 right-4 z-10 space-x-2">
        <Button onClick={onSave}>Save Workflow</Button>
      </div>
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
