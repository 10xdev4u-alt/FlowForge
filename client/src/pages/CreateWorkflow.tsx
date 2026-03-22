import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../api/axios';
import { Button } from '../components/ui/Button';
import { Input } from '../components/ui/Input';

export const CreateWorkflow = () => {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const navigate = useNavigate();

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const { data } = await api.post('/workflows', { name, description });
      navigate(`/editor/${data.id}`);
    } catch (error) {
      alert('Failed to create workflow');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <form onSubmit={handleCreate} className="p-8 bg-white shadow-lg rounded-lg w-96 space-y-4">
        <h1 className="text-2xl font-bold text-center">New Workflow</h1>
        <Input placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} required />
        <Input placeholder="Description" value={description} onChange={(e) => setDescription(e.target.value)} />
        <Button type="submit" className="w-full">Create and Edit</Button>
      </form>
    </div>
  );
};
