import React from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';

type GithubProfileProps = {
  user: {
    Bio?: string;
    avatar_url?: string;
    blog?: string;
    collaborators?: number;
    company?: string;
    email?: string;
    evaluation?: string;
    followers?: number;
    following?: number;
    location?: string;
    login_name?: string;
    name?: string;
    nationality?: string;
    public_repos?: number;
    score?: number;
    total_private_repos?: number;
  };
};

const UserProfile: React.FC<GithubProfileProps> = ({ user }) => {
  return (
    <Card className="max-w-md my-auto bg-white shadow-md rounded-lg border border-gray-200 m-2">
      <CardHeader className="flex flex-col items-center py-6">
        <Avatar className="w-24 h-24 rounded-full">
          <AvatarImage src={user.avatar_url} alt={user.avatar_url} />
          <AvatarFallback className="rounded-lg">CN</AvatarFallback>
        </Avatar>
        <CardTitle className="text-xl font-semibold mt-4">{user.name}</CardTitle>
        <p className="text-gray-500">@{user.login_name}</p>
      </CardHeader>

      <CardContent className="px-6 pb-6">
        <p className="text-gray-800 mb-2">
          <strong>Bio:</strong> {user.Bio || 'N/A'}
        </p>
        <p className="text-gray-800">
          <strong>Company:</strong> {user.company || 'N/A'}
        </p>
        <p className="text-gray-800">
          <strong>Location:</strong> {user.location || 'N/A'}
        </p>
        <p className="text-gray-800">
          <strong>Nationality:</strong> {user.nationality || 'N/A'}
        </p>

        <div className="grid grid-cols-2 gap-4 mt-4">
          <p className="text-gray-800">
            <strong>Followers:</strong> {user.followers}
          </p>
          <p className="text-gray-800">
            <strong>Following:</strong> {user.following}
          </p>
          <p className="text-gray-800">
            <strong>Public Repos:</strong> {user.public_repos}
          </p>
          <p className="text-gray-800">
            <strong>Private Repos:</strong> {user.total_private_repos}
          </p>
          <p className="text-gray-800">
            <strong>Collaborators:</strong> {user.collaborators}
          </p>
          <p className="text-gray-800">
            <strong>Score:</strong> {user?.score?.toFixed(1)}
          </p>
        </div>

        <p className="text-gray-800 mt-4">
          <strong>Email:</strong> {user.email || 'UnKnown'}
        </p>
        <p className="text-gray-800">
          <strong>Blog:</strong> {user.blog || 'UnKnown'}
        </p>
        {/* <p className="text-gray-800"><strong>Evaluation:</strong> {user.evaluation || "N/A"}</p> */}
      </CardContent>
    </Card>
  );
};

export default UserProfile;
