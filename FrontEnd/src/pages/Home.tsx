import { AppSidebar } from '@/components/app-sidebar';
import Rank from '@/components/common/Rank';
import SearchBar from '@/components/common/SearchBar';
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from '@/components/ui/breadcrumb';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Separator } from '@/components/ui/separator';
import { SidebarInset, SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar';

import { get } from '@/fetch';
import { Ranks, ResponseLogin, ResponseRank, ResponseUserInfo, UserInfo } from '@/type';
import { useEffect, useState } from 'react';

function Home() {
  const [rank, setRank] = useState<Ranks[]>();
  const [domain, setDomain] = useState<string[]>();
  useEffect(() => {
    handleGithub();
    handleRank();
    handleUser();
    handleDomain();
  }, []);
  const [user, setUser] = useState<UserInfo>({
    avatar_url: '',
    email: '',
    name: '',
    score: 0,
  });
  const handleUser = () => {
    get<ResponseUserInfo>('/api/v1/user/getInfo', true).then((res) => {
      setUser(res.data.user);
    });
  };
  const handleDomain = () => {
    get<ResponseUserInfo>('/api/v1/user/getInfo', true).then((res) => {
      setDomain(res.data.domain);
    });
  };
  const handleRank = () => {
    get<ResponseRank>('/api/v1/user/getRank', true).then((res) => {
      setRank(res.data.leaderboard);
    });
  };
  const handleGithub = () => {
    const code = new URLSearchParams(window.location.search).get('code') || '';
    console.log('code', code);
    get<ResponseLogin>('/api/v1/auth/callBack?code=' + code, false).then((res) => {
      localStorage.setItem('token', res?.data.token || '');
      console.log('res', res);
    });
  };
  return (
    <>
      <SidebarProvider>
        <AppSidebar />
        <SidebarInset>
          <header className="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12">
            <div className="flex items-center gap-2 px-4">
              <SidebarTrigger className="-ml-1" />
              <Separator orientation="vertical" className="mr-2 h-4" />
              <Breadcrumb>
                <BreadcrumbList>
                  <BreadcrumbItem className="hidden md:block">
                    <BreadcrumbLink href="#">GitEval</BreadcrumbLink>
                  </BreadcrumbItem>
                  <BreadcrumbSeparator className="hidden md:block" />
                  <BreadcrumbItem>
                    <BreadcrumbPage>Data Fetching</BreadcrumbPage>
                  </BreadcrumbItem>
                </BreadcrumbList>
              </Breadcrumb>
            </div>
          </header>
          <div className="flex flex-1 flex-col gap-4 p-4 pt-0">
            <div className="min-h-[100vh] flex-1 rounded-xl bg-muted/50 md:min-h-min flex">
              <div>
                <Card className="m-4 w-[50vw]">
                  <CardHeader>
                    <CardTitle>根据大模型预测</CardTitle>
                  </CardHeader>
                  <CardContent>
                    <p>领域:{domain}</p>
                  </CardContent>
                  <CardFooter>
                    <p>地区{user.nationality}</p>
                  </CardFooter>
                </Card>
                <Rank data={rank}></Rank>
              </div>
              <SearchBar></SearchBar>
            </div>
          </div>
        </SidebarInset>
      </SidebarProvider>
    </>
  );
}
export default Home;
