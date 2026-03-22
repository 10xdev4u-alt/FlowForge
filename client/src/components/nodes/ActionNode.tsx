import React from 'react';
import { Handle, Position } from 'reactflow';

export const ActionNode = ({ data }: { data: any }) => {
  return (
    <div className="px-4 py-2 shadow-md rounded-md bg-white border-2 border-blue-500">
      <Handle type="target" position={Position.Top} className="!w-2 !h-2 !bg-blue-500" />
      <div className="flex flex-col">
        <div className="text-xs font-bold text-blue-500">ACTION</div>
        <div className="text-sm">{data.label}</div>
      </div>
      <Handle type="source" position={Position.Bottom} className="!w-2 !h-2 !bg-blue-500" />
    </div>
  );
};
