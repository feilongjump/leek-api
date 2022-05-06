<?php

namespace App\Observers;

use App\Models\Article;

class ArticleObserver
{
    public function creating(Article $article)
    {
        $article->user_id = auth()->id() ?? 1;
    }
}
