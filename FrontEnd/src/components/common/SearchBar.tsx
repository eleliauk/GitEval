import { get } from '@/fetch';
import { ModelUser, ResponseSearchs } from '@/type';
import { useState } from 'react';

export default function SearchBar() {
  const [nation, setNation] = useState('');
  const [domain, setDomain] = useState('');
  const [results, setResults] = useState<ModelUser[]>([]);
  const [currentPage, setCurrentPage] = useState(1); // 当前页码
  const [totalPages, setTotalPages] = useState(1); // 总页数
  const country = [
    '中国',
    '美国',
    '英国',
    '日本',
    '法国',
    '德国',
    '意大利',
    '加拿大',
    '澳大利亚',
    '俄罗斯',
    '印度',
    '巴西',
    '西班牙',
    '墨西哥',
    '波兰',
    '韩国',
    '瑞典',
  ];
  const field = [
    '前端开发',
    '后端开发',
    '数据科学',
    '人工智能',
    '移动开发',
    '区块链',
    '网络安全',
    '游戏开发',
    '数据库开发',
    '云计算',
  ];

  const handleSearch = (page = 1) => {
    get<ResponseSearchs>(
      `/api/v1/user/search?nation=${nation}&domain=${domain}&page=${page}&page_size=10`,
      true,
    ).then((res) => {
      setResults(res.data.users);
      setTotalPages(res.data.users.length);
    });
  };

  const handlePageChange = (newPage: number) => {
    setCurrentPage(newPage);
    handleSearch(newPage);
  };

  return (
    <div className="p-6 min-w-[25vw] mx-auto">
      <div className="bg-white shadow-md rounded-lg p-6 space-y-4">
        {/* 国家输入框 */}
        <div>
          <label htmlFor="country" className="block text-sm font-medium text-gray-700">
            选择国家
          </label>
          <select
            id="country"
            value={nation}
            onChange={(e) => setNation(e.target.value)}
            className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:ring-indigo-500 focus:border-indigo-500"
          >
            <option value="" disabled>
              选择一个国家
            </option>
            {country.map((f) => (
              <option key={f} value={f}>
                {f}
              </option>
            ))}
          </select>
        </div>
        <div>
          <label htmlFor="field" className="block text-sm font-medium text-gray-700">
            选择领域
          </label>
          <select
            id="field"
            value={domain}
            onChange={(e) => setDomain(e.target.value)}
            className="mt-1 p-2 border border-gray-300 rounded-md w-full focus:ring-indigo-500 focus:border-indigo-500"
          >
            <option value="" disabled>
              选择一个领域
            </option>
            {field.map((f) => (
              <option key={f} value={f}>
                {f}
              </option>
            ))}
          </select>
        </div>
        <button
          onClick={() => handleSearch(1)}
          className="w-full bg-slate-900 text-white py-2 rounded-md hover:bg-slate-500 focus:bg-slate-500"
        >
          搜索
        </button>
      </div>
      <div className="mt-6 bg-white shadow-md rounded-lg p-6">
        <h3 className="text-lg font-medium text-gray-900 mb-4">搜索结果</h3>
        {results.length > 0 ? (
          <ul className="space-y-2">
            {results.map((result, index) => (
              <li key={index} className="text-gray-700 w-[20vw]">
                <div key={index} className="flex items-center space-x-2 hover:bg-gray-100 w-[20vw]">
                  <img
                    src={result.avatar_url}
                    alt={result.name}
                    className="w-10 h-10 rounded-full"
                  />
                  <span>{result.name}</span>
                  <span>{result.score?.toFixed(1)}</span>
                </div>
              </li>
            ))}
          </ul>
        ) : (
          <p className="text-gray-500">暂无结果</p>
        )}
        {/* 翻页功能 */}
        <div className="flex justify-between mt-4">
          <button
            onClick={() => handlePageChange(currentPage - 1)}
            disabled={currentPage === 1}
            className="bg-gray-300 text-gray-700 py-1 px-3 rounded-md hover:bg-gray-400 disabled:bg-gray-200"
          >
            上一页
          </button>
          {/* <span className="text-gray-700">
            第 {currentPage} 页，共 {totalPages} 页
          </span> */}
          <button
            onClick={() => handlePageChange(currentPage + 1)}
            disabled={currentPage === totalPages}
            className="bg-gray-300 text-gray-700 py-1 px-3 rounded-md hover:bg-gray-400 disabled:bg-gray-200"
          >
            下一页
          </button>
        </div>
      </div>
    </div>
  );
}
