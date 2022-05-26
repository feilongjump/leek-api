<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\SoftDeletes;

class ProjectColumnCard extends Model
{
    use HasFactory, SoftDeletes;

    /**
     * The attributes that are mass assignable.
     *
     * @var string[]
     */
    protected $fillable = [
        'project_column_id', 'name',
    ];

    public function content()
    {
        return $this->morphOne(Content::class, 'contentable');
    }

    public function column()
    {
        return $this->belongsTo(ProjectColumn::class, 'project_column_id');
    }
}
