<?php

namespace App\Observers;

use App\Models\Article;
use Illuminate\Support\Arr;

class ArticleObserver
{
    public function creating(Article $article)
    {
        $article->user_id = auth()->id() ?? 1;
    }

    public function created(Article $article)
    {
        $this->saveContent($article);
    }

    public function saving(Article $article)
    {
        if ($article->isClean()) $article[$article::UPDATED_AT] = now();
    }

    public function saved(Article $article)
    {
        $this->saveContent($article);
    }

    private function saveContent(Article $article)
    {
        $type = request('type', 'markdown');

        $data = Arr::only(request('content', []), $type);

        $article->content()->updateOrCreate(['contentable_id' => $article->id], $data);

        $article->loadMissing('content');
    }
}
