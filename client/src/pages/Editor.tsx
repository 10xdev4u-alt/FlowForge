import React from 'react';
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

import { Node } from 'reactflow';
import { Input } from '../components/ui/Input';

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
  const [selectedNode, setSelectedNode] = React.useState<Node | null>(null);

  React.useEffect(() => {
    const fetchCanvas = async () => {
      try {
        const { data } = await api.get(`/workflows/${id}/canvas`);
        if (data && data.nodes && data.edges) {
          const parseData = (val: string) => {
             try {
               return JSON.parse(val);
             } catch {
               return JSON.parse(atob(val));
             }
          }
          setNodes(parseData(data.nodes));
          setEdges(parseData(data.edges));
        }
      } catch (error) {
        setNodes([{ id: '1', type: 'trigger', position: { x: 100, y: 100 }, data: { label: 'Webhook' } }]);
        setEdges([]);
      }
    };
    fetchCanvas();
  }, [id, setNodes, setEdges]);

  const onSave = async () => {
    try {
      await api.post(`/workflows/${id}/canvas`, { nodes, edges });
      alert('Workflow saved!');
    } catch (error) {
      alert('Failed to save workflow');
    }
  };

  const onNodeClick = (_: React.MouseEvent, node: Node) => {
    setSelectedNode(node);
  };

  return (
    <div className="w-full h-screen relative flex">
      <div className="flex-1 relative">
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
          onNodeClick={onNodeClick}
        >
          <Controls />
          <Background />
        </ReactFlow>
      </div>
      {selectedNode && (
        <div className="w-80 bg-white border-l p-4 shadow-xl z-20 overflow-y-auto">
          <div className="flex justify-between items-center mb-6">
            <h2 className="text-lg font-bold">Node Settings</h2>
            <button className="text-gray-500 hover:text-black" onClick={() => setSelectedNode(null)}>✕</button>
          </div>
          <div className="space-y-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Type: <span className="uppercase font-bold text-xs bg-gray-100 px-2 py-1 rounded">{selectedNode.type}</span></label>
              <p className="text-xs text-gray-500">ID: {selectedNode.id}</p>
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">Display Name</label>
              <Input 
                value={selectedNode.data.label || ''} 
                onChange={(e) => {
                  const newNodes = nodes.map(n => n.id === selectedNode.id ? { ...n, data: { ...n.data, label: e.target.value } } : n);
                  setNodes(newNodes);
                }} 
              />
            </div>
            {/* Node specific settings could go here */}
          </div>
        </div>
      )}
    </div>
  );
};
