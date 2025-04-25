import React, { useState, useEffect } from 'react';
import ReactMarkdown from 'react-markdown';
import { Card, CardContent, CardHeader } from '../ui/card';

interface AnimatedMarkdownProps {
  content: string;
  speed?: number;
}

const AnimatedMarkdown: React.FC<AnimatedMarkdownProps> = ({ content, speed = 50 }) => {
  const [displayText, setDisplayText] = useState('');
  const [index, setIndex] = useState(0);

  const formattedContent = content.replace(/\n/g, '  \n');

  useEffect(() => {
    if (index < formattedContent.length) {
      const timer = setTimeout(() => {
        setDisplayText((prev) => prev + formattedContent[index]);
        setIndex(index + 1);
      }, speed);

      return () => clearTimeout(timer);
    }
  }, [index, formattedContent, speed]);

  return (
    <div className="p-4 text-gray-700 w-[40vw]">
      <Card className="w-[40vw] min-h-[70vh] my-auto bg-white shadow-md rounded-lg border border-gray-200">
        <CardHeader className="flex flex-col items-center py-6"></CardHeader>
        <CardContent className="px-6 pb-6">
          <ReactMarkdown>{displayText}</ReactMarkdown>
        </CardContent>
      </Card>
    </div>
  );
};

export default AnimatedMarkdown;
