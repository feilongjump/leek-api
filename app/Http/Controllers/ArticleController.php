<?php

namespace App\Http\Controllers;

use App\Http\Requests\ArticleRequest;
use App\Http\Resources\ArticleResource;
use App\Models\Article;
use Illuminate\Http\Request;

class ArticleController extends Controller
{
    public function __construct()
    {
        $this->middleware(['auth:sanctum']);
    }

    public function index(Request $request)
    {
        $articles = Article::whereUserId(auth()->user()->id)
            ->latest()
            ->paginate($request->get('per_page', 10));

        return ArticleResource::collection($articles);
    }

    public function store(ArticleRequest $request)
    {
        $this->authorize('create', Article::class);

        return new ArticleResource(Article::create($request->all()));
    }

    public function show(Article $article)
    {
        $this->authorize('view', $article);

        $article->loadMissing('content');

        return new ArticleResource($article);
    }

    public function update(ArticleRequest $request, Article $article)
    {
        $this->authorize('update', $article);

        $article->update($request->all());

        return new ArticleResource($article);
    }

    public function destroy(Article $article)
    {
        $this->authorize('delete', $article);

        $article->delete();

        return $this->withNoContent();
    }
}
