import React from 'react'

interface MessageProps {
  type: 'success' | 'error';
  text: string;
}

const Message: React.FC<MessageProps> = ({ type, text }) => {
  const baseStyle =
    'px-4 py-3 rounded-md text-sm font-medium shadow-md transition-all duration-300 mb-4';
  const styles = {
    success: 'bg-green-100 text-green-800 border border-green-300',
    error: 'bg-red-100 text-red-800 border border-red-300',
  };

  return (
    <div className={`${baseStyle} ${styles[type]}`}>
      {text}
    </div>
  );
};

export default Message;