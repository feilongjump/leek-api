<?php

namespace App\Observers;

use App\Models\ProjectColumnCard;
use Illuminate\Support\Arr;

class ProjectColumnCardObserver
{
    public function created(ProjectColumnCard $card)
    {
        $this->saveContent($card);
    }

    public function saving(ProjectColumnCard $card)
    {
        if ($card->isClean()) $card[$card::UPDATED_AT] = now();
    }

    public function saved(ProjectColumnCard $card)
    {
        $this->saveContent($card);
    }

    private function saveContent(ProjectColumnCard $card)
    {
        $type = request('type', 'markdown');

        $data = Arr::only(request('content', []), $type);

        $card->content()->updateOrCreate(['contentable_id' => $card->id], $data);

        $card->loadMissing('content');
    }
}
