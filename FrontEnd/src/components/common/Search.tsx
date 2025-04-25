import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

export function Search() {
  return (
    <div className="flex  w-[20vw] items-center space-x-2 justify-center">
      <Input type="text" placeholder="Search" />
      <Button type="submit">Search</Button>
    </div>
  );
}
