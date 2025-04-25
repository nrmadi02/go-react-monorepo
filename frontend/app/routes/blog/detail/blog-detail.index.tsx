import axios from "axios";
import type { Route } from "./+types/blog-detail.index";

export function meta({ data }: Route.MetaArgs) {
  return [
    { title: data.title },
    { name: "description", content: data.content },
  ];
}

interface Post {
  id: number;
  title: string;
  content: string;
}


export async function loader({ params }: Route.LoaderArgs) {
  const post = await axios.get<{
    data: Post
  }>(`http://127.0.0.1:8000/post/${params.id}`);

  return post.data.data;
}

export default function BlogDetail(props: Route.ComponentProps ) {
  return (
    <div className="p-5">
      <div className="mt-3 border p-3">
        <h1 className="text-3xl">{props.loaderData.title}</h1>
        <p className="mt-3">{props.loaderData.content}</p>
      </div>
    </div>
  );
}
