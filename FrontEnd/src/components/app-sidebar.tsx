import * as React from 'react';
import { BookOpen, SquareTerminal } from 'lucide-react';

import { NavMain } from '@/components/nav-main';
//import { NavProjects } from '@/components/nav-projects';
import { NavUser } from '@/components/nav-user';
//import { TeamSwitcher } from "@/components/team-switcher"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from '@/components/ui/sidebar';
const data = {
  user: {
    name: 'shadcn',
    email: 'm@example.com',
    avatar: '/avatars/shadcn.jpg',
  },
  navMain: [
    {
      title: 'Playground',
      url: '/Home',
      icon: SquareTerminal,
      items: [
        {
          title: 'Home',
          url: '/Home',
        },
      ],
    },
    {
      title: 'Documentation',
      url: '/Eval',
      icon: BookOpen,
      items: [
        {
          title: 'Evaluation',
          url: '/Eval',
        },
      ],
    },
  ],
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader></SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}
