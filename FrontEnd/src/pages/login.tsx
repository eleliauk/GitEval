import GithubIcon from '@/components/common/githubIcon';
import { Button } from '@/components/ui/button';
import { Card, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
// import { get } from '@/fetch';
// import { ResponseLogin } from '@/type';

function Login() {
  function test() {
    window.location.href = 'http://47.92.102.209:8080/api/v1/auth/login';
  }
  return (
    <div className="flex items-center justify-center w-full min-h-screen p-6">
      <div className="grid gap-6 w-full max-w-sm">
        <h1 className="text-3xl font-bold">登录</h1>
        <Card>
          <CardHeader>
            <CardTitle>GitEval</CardTitle>
            <CardDescription>超级无敌的github排行榜</CardDescription>
          </CardHeader>
          <CardFooter>
            <Button className="w-full inline-flex gap-4" onClick={test}>
              <GithubIcon className="h-4 w-4" />
              GitHub
            </Button>
          </CardFooter>
        </Card>
      </div>
    </div>
  );
}
export default Login;
