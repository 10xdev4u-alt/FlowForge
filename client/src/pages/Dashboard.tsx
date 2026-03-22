import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../api/axios';
import { Button } from '../components/ui/Button';

interface Workflow {
  id: string;
  name: string;
  description: { String: string; Valid: boolean };
  is_active: boolean;
}

export const Dashboard = () => {
  const [workflows, setWorkflows] = useState<Workflow[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    fetchWorkflows();
  }, []);

  const fetchWorkflows = async () => {
    try {
      const { data } = await api.get('/workflows');
      setWorkflows(data || []);
    } catch (error) {
      console.error('Failed to fetch workflows');
    }
  };

  return (
    <div className="p-8 max-w-6xl mx-auto space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-3xl font-bold">My Workflows</h1>
        <Button onClick={() => navigate('/workflows/new')}>Create Workflow</Button>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {workflows.map((wf) => (
          <div key={wf.id} className="p-4 border rounded shadow hover:shadow-md transition-shadow cursor-pointer" onClick={() => navigate(`/editor/${wf.id}`)}>
            <h2 className="text-xl font-semibold">{wf.name}</h2>
            <p className="text-gray-600">{wf.description.String}</p>
            <div className="mt-4 flex items-center space-x-2">
              <span className={`w-3 h-3 rounded-full ${wf.is_active ? 'bg-green-500' : 'bg-gray-400'}`}></span>
              <span className="text-sm">{wf.is_active ? 'Active' : 'Inactive'}</span>
            </div>
          </div>
        ))}
        {workflows.length === 0 && (
          <div className="col-span-full text-center py-20 text-gray-500">
            No workflows found. Create your first one!
          </div>
        )}
      </div>
    </div>
  );
};
