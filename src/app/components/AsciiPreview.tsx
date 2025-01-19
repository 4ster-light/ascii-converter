import React from 'react';

interface AsciiPreviewProps {
  ascii: string;
}

export const AsciiPreview: React.FC<AsciiPreviewProps> = ({ ascii }) => {
  return (
    <div className="card bg-base-300 duration-200">
      <div className="card-body p-4">
        <div className="flex justify-center">
          <pre
            className="font-mono text-xs md:text-sm lg:text-base whitespace-pre text-base-content"
            dangerouslySetInnerHTML={{ __html: ascii }}
            style={{
              lineHeight: '0.5',
              letterSpacing: '0',
              textAlign: 'center',
              fontFamily: 'Consolas, Monaco, "Courier New", monospace',
              fontSize: '2px',
              transform: 'scale(1.2)',
              transformOrigin: 'center',
              margin: '10px 0',
              display: 'inline-block',
              maxWidth: '100%',
              overflow: 'hidden'
            }}
          />
        </div>
      </div>
    </div>
  );
}
