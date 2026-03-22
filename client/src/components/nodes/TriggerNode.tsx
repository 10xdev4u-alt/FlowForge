import React from 'react';
import { Handle, Position } from 'reactflow';

export const TriggerNode = ({ data }: { data: any }) => {
  return (
    <div className="px-4 py-2 shadow-md rounded-md bg-white border-2 border-green-500">
      <div className="flex flex-col">
        <div className="text-xs font-bold text-green-500">TRIGGER</div>
        <div className="text-sm">{data.label || 'Webhook'}</div>
      </div>
      <Handle type="source" position={Position.Bottom} className="!w-2 !h-2 !bg-green-500" />
    </div>
  );
};
