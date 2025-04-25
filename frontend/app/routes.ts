import { type RouteConfig, index, layout, route } from "@react-router/dev/routes";

export default [
  index("routes/home.tsx"),
  layout("routes/blog/blog.layout.tsx", [
    route("blog", "routes/blog/blog.index.tsx"),
    route("blog/detail/:id", "routes/blog/detail/blog-detail.index.tsx"),
  ]),
] satisfies RouteConfig;
