<?php

namespace App\Observers;

use App\Models\ProjectColumn;

class ProjectColumnObserver
{
    public function deleted(ProjectColumn $column)
    {
        $column->cards()->delete();
    }
}
