import axios from "axios";
import type { Route } from "./+types/blog.index";
import { Link } from "react-router";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Blog" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

interface Post {
  id: number;
  title: string;
  content: string;
}

export async function loader() {
  const posts = await axios.get<{
    data: Post[];
  }>(`http://127.0.0.1:8000/post`);
  return posts.data;
}

export default function Blog(props: Route.ComponentProps) {
  const posts = props.loaderData;

  return (
    <div className="p-5">
      <h1 className="text-3xl">Blog</h1>
      <div className="mt-3">
        {posts.data.map((post) => (
          <Link key={post.id} to={`/blog/detail/${post.id}`}>
            <div className="mt-3 p-3 border hover:opacity-50 transition-all cursor-pointer">
              <h2 className="text-2xl">{post.title}</h2>
              <p className="mt-2">{post.content}</p>
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
}
