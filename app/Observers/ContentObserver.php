<?php

namespace App\Observers;

use App\Models\Content;

class ContentObserver
{
    public function saving(Content $content)
    {
        $this->toHtml($content);
    }

    private function toHtml(Content $content)
    {
        if ($content->isDirty('markdown') && !empty($content->markdown)) {
            $content->body = (new \Parsedown())->text($content->markdown);
        }
    }
}
